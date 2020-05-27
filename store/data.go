package store

import (
	shell "github.com/ipfs/go-ipfs-api"
)

type ipfs struct {
	shell *shell.Shell
}

func newDataStore(args string) DataOperator {
	return &ipfs{
		shell.NewShell(args),
	}
}

// Put adds the interface to IPFS and returns the corresponding content identifier (CID)
func (i ipfs) Put(value []byte) (string, error) {
	cid, err := i.shell.DagPut(value, "json", "cbor")
	if err != nil {
		return "", err
	}
	return cid, nil
}

// Get returns the corresponding data based on the key
func (i ipfs) Get(key []byte) ([]byte, error) {
	panic("implement me")
}

func (i ipfs) Has(key []byte) (bool, error) {
	panic("implement me")
}

func (i ipfs) Delete(key []byte) error {
	panic("implement me")
}

func (i ipfs) Close() error {
	panic("implement me")
}
