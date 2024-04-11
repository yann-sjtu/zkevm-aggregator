package state

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	// MockL1InfoRootHex is used to send batches to the Executor
	// the number below represents this formula:
	//
	// 	mockL1InfoRoot := common.Hash{}
	// for i := 0; i < len(mockL1InfoRoot); i++ {
	// 	  mockL1InfoRoot[i] = byte(i)
	// }
	MockL1InfoRootHex = "0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
)

// Batch struct
type Batch struct {
	BatchNumber     uint64
	Coinbase        common.Address
	BatchL2Data     []byte
	StateRoot       common.Hash
	LocalExitRoot   common.Hash
	AccInputHash    common.Hash
	L1InfoTreeIndex uint32
	// Timestamp (<=incaberry) -> batch time
	// 			 (>incaberry) -> minTimestamp used in batch creation, real timestamp is in virtual_batch.batch_timestamp
	Timestamp      time.Time
	Transactions   []types.Transaction
	GlobalExitRoot common.Hash
	ForcedBatchNum *uint64
}

// Sequence represents the sequence interval
type Sequence struct {
	FromBatchNumber uint64
	ToBatchNumber   uint64
}

var mockL1InfoRoot = common.HexToHash(MockL1InfoRootHex)

// GetMockL1InfoRoot returns an instance of common.Hash set
// with the value provided by the const MockL1InfoRootHex
func GetMockL1InfoRoot() common.Hash {
	return mockL1InfoRoot
}
