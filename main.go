package main

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/m8b-dev/spike-deposit-2-genesis/contract/binding"
	"github.com/schollz/progressbar/v3"
	"math/big"
	"os"
	"sync"
	"time"
)

type JSONData struct {
	Pubkey                string `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	Amount                uint64 `json:"amount"`
	Signature             string `json:"signature"`
	DepositDataRoot       string `json:"deposit_data_root"`
}

const address = "0x00000000219ab540356cbb839cbe05303d7705fa"
const infuraUrl = ""
const startBlk = uint64(12775113)
const endBlk = uint64(12975113)
const findEndBlock = false

func main() {
	eth, err := ethclient.Dial(infuraUrl)
	if err != nil {
		panic(err)
	}
	output := make([]JSONData, 0)
	addr := common.HexToAddress(address)
	deposits, err := binding.NewBinding(addr, eth)
	if err != nil {
		panic(err)
	}
	abi, err := binding.BindingMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	maxBlk := endBlk
	if findEndBlock {
		maxBlk, err = eth.BlockNumber(context.Background())
		if err != nil {
			panic(err)
		}
	}

	blockData := multiThreadedFetch(eth, startBlk, maxBlk, addr)
	for _, blk := range blockData {
		for _, txData := range blk.data {
			data := txData.input
			// cut off first 4 bytes of method identifier
			data = data[4:]
			params, err := abi.Methods["deposit"].Inputs.Unpack(data)
			if err != nil {
				panic(err)
			}
			dataRoot, ok := params[3].([32]byte)
			if !ok {
				panic("got invalid type for data root")
			}

			var depEvent *binding.BindingDepositEvent
			for _, evnt := range txData.logs {
				depEvent, err = deposits.ParseDepositEvent(*evnt)
				if err == nil {
					break
				}
			}
			if err != nil {
				panic(err)
			}
			//  solc: bytes memory amount = to_little_endian_64(uint64(deposit_amount));

			output = append(output, JSONData{
				Pubkey:                hexutil.Encode(depEvent.Pubkey),
				WithdrawalCredentials: hexutil.Encode(depEvent.WithdrawalCredentials),
				Amount:                binary.LittleEndian.Uint64(depEvent.Amount),
				Signature:             hexutil.Encode(depEvent.Signature),
				DepositDataRoot:       hexutil.Encode(dataRoot[:]),
			})
		}
	}

	outputMarshaled, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("failed to marshal data, data: %+v", output)
		panic(err)
	}
	err = os.WriteFile("./deposit_data.json", outputMarshaled, 0600)
	if err != nil {
		fmt.Printf("failed to write to file, data: %+v", outputMarshaled)
		panic(err)
	}
	fmt.Printf("Scan done!\n Deposit data written OK\n Found %d deposits", len(output))
}

type mutexedUint struct {
	val uint
	mut sync.Mutex
}

func (m *mutexedUint) Add(x uint) {
	m.mut.Lock()
	m.val += x
	m.mut.Unlock()
}
func (m *mutexedUint) Sub(x uint) {
	m.mut.Lock()
	m.val -= x
	m.mut.Unlock()
}
func (m *mutexedUint) Get() uint {
	m.mut.Lock()
	x := m.val
	m.mut.Unlock()
	return x
}

func multiThreadedFetch(client *ethclient.Client, from, to uint64, filter common.Address) []fetchBlockOutput {
	const maxThreads = 80
	runs := int(to - from)
	activeThreads := mutexedUint{val: 0, mut: sync.Mutex{}}

	bar := progressbar.Default(int64(runs), "scanning blocks...")

	output := make([]fetchBlockOutput, runs)
	for i := 0; i < runs; i++ {
		for activeThreads.Get() >= maxThreads {
			time.Sleep(time.Millisecond)
		}
		activeThreads.Add(1)
		go func(i int) {
			output[i] = fetchBlock(client, from+uint64(i), filter)
			activeThreads.Sub(1)
			// dont let ui break the process
			_ = bar.Add(1)
		}(i)
	}
	for activeThreads.Get() > 0 {
		time.Sleep(time.Millisecond)
	}
	_ = bar.Finish()
	_ = bar.Close()
	return output
}

type fetchBlockOutput struct {
	block uint64
	data  []fetchBlockEntry
}
type fetchBlockEntry struct {
	hash  common.Hash
	input []byte
	logs  []*types.Log
}

func fetchBlock(client *ethclient.Client, block uint64, filter common.Address) fetchBlockOutput {
	err := errors.New("fake err")
	var blk *types.Block
	// try until success (tmp network issues etc)
	for err != nil {
		blk, err = client.BlockByNumber(context.Background(), big.NewInt(0).SetUint64(block))
	}
	txns := make([]*types.Transaction, 0)
	output := fetchBlockOutput{
		block: block,
		data:  make([]fetchBlockEntry, 0),
	}
	for _, txn := range blk.Transactions() {
		if txn.To() != nil && *txn.To() == filter {
			txns = append(txns, txn)
		}
	}
	for _, txn := range txns {
		err = errors.New("fake err")
		var rcpt *types.Receipt
		// try until success (tmp network issues etc)
		for err != nil {
			rcpt, err = client.TransactionReceipt(context.Background(), txn.Hash())
		}
		if rcpt.Status == 1 {
			output.data = append(output.data, fetchBlockEntry{
				hash:  txn.Hash(),
				input: txn.Data(),
				logs:  rcpt.Logs,
			})
		}
	}
	return output
}
