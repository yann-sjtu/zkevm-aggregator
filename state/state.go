package state

import (
	"context"
	"sync"

	"github.com/0xPolygonHermez/zkevm-aggregator/event"
	"github.com/0xPolygonHermez/zkevm-aggregator/state/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
)

var (
	// ZeroHash is the hash 0x0000000000000000000000000000000000000000000000000000000000000000
	ZeroHash = common.Hash{}
	// ZeroAddress is the address 0x0000000000000000000000000000000000000000
	ZeroAddress = common.Address{}
)

// State is an implementation of the state
type State struct {
	cfg Config
	storage
	eventLog *event.EventLog
}

// NewState creates a new State
func NewState(cfg Config, storage storage, eventLog *event.EventLog) *State {
	var once sync.Once
	once.Do(func() {
		metrics.Register()
	})

	state := &State{
		cfg:      cfg,
		storage:  storage,
		eventLog: eventLog,
	}

	return state
}

// BeginStateTransaction starts a state transaction
func (s *State) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := s.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
