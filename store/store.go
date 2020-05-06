package store

import (
	"github.com/floydeconomy/arisaedo-go/store/chain"
	"github.com/floydeconomy/arisaedo-go/store/data"
)

type Store struct {
	db     data.Operator
	client chain.Operator
}

type Options struct {
	Db    string
	Chain string
}

func New(o Options) *Store {
	return &Store{
		db:     newDataStore(o.Db),
		client: newChainStore(o.Chain),
	}
}
