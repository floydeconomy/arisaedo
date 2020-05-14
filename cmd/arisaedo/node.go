package main

import (
	"context"
	"github.com/floydeconomy/arisaedo-go/co"
	"github.com/floydeconomy/arisaedo-go/store"
)

// todo: implement logDB
type Node struct {
	goes co.Goes

	store *store.Store
}

func New(store *store.Store) *Node {
	return &Node{
		store: store,
	}
}

func (n *Node) Run(ctx context.Context) error {
	// todo: implement houseKeeping and p2p comm.sync
	return nil
}
