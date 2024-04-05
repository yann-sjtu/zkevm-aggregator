package merkletree

import (
	"context"
	"time"

	"github.com/0xPolygonHermez/zkevm-aggregator/log"
	"github.com/0xPolygonHermez/zkevm-aggregator/merkletree/hashdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewMTDBServiceClient creates a new MTDB client.
func NewMTDBServiceClient(ctx context.Context, c Config) (hashdb.HashDBServiceClient, *grpc.ClientConn, context.CancelFunc) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	const maxWaitSeconds = 120
	_, cancel := context.WithTimeout(ctx, maxWaitSeconds*time.Second)

	log.Infof("trying to connect to merkletree: %v", c.URI)
	mtDBConn, err := grpc.NewClient(c.URI, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Infof("connected to merkletree")

	mtDBClient := hashdb.NewHashDBServiceClient(mtDBConn)
	return mtDBClient, mtDBConn, cancel
}
