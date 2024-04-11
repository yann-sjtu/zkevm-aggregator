package state

import (
	"github.com/0xPolygonHermez/zkevm-aggregator/db"
)

// Config is state config
type Config struct {
	// ChainID is the L2 ChainID provided by the Network Config
	ChainID uint64
	// DB is the database configuration
	DB db.Config `mapstructure:"DB"`
}
