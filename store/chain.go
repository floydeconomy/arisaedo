package store

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type eth struct {
	client *ethclient.Client
}

// todo: fix log.fatal
func newChainStore(args string) ChainOperator {
	c, err := ethclient.Dial(args) // localhost:8545 for ganache-cli
	if err != nil {
		log.Fatal("ethclient failed")
	}
	return &eth{c}
}

func (e eth) Get(key []byte) ([]byte, error) {
	panic("implement me")
}

func (e eth) Has(key []byte) (bool, error) {
	panic("implement me")
}

func (e eth) Put(value []byte) (string, error) {
	panic("implement me")
}

func (e eth) Delete(key []byte) error {
	panic("implement me")
}

func (e eth) Close() error {
	panic("implement me")
}
