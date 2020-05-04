package node

import (
	"context"
	"github.com/floydeconomy/arisaedo-go/co"
	"github.com/floydeconomy/arisaedo-go/internal"
)

// todo: implement logDB
type Node struct {
	goes co.Goes

	repo *internal.Repository
}

func New() *Node {
	return &Node{
		repo: internal.New(),
	}
}

func (n *Node) Run(ctx context.Context) error {
	// todo: implement houseKeeping and p2p comm.sync
	return nil
}
