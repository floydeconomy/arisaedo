package store

import (
	"encoding/json"
	eth "github.com/ethereum/go-ethereum/ethclient"
	"github.com/floydeconomy/arisaedo-go/common"
	"github.com/floydeconomy/arisaedo-go/kv"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/pkg/errors"
)

type Store struct {
	data kv.Store

	shell  *shell.Shell
	client *eth.Client
}

type Options struct {
	Db    string
	Chain string
}

func New(o Options) (*Store, error) {
	s := shell.NewShell(o.Db)
	c, err := eth.Dial(o.Chain) // localhost:8545 for ganache-cli
	if err != nil {
		return nil, errors.Wrapf(err, "eth db failed at [%v]", o.Chain)
	}
	return &Store{
		shell:  s,
		client: c,
	}, nil
}

// Put adds the interface to IPFS and returns the corresponding content identifier (CID)
func (s Store) Put(x interface{}) (*common.Identifier, error) {
	// marshall json
	m, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	// ipfs put
	cid, err := s.shell.DagPut(m, "json", "cbor")
	if err != nil {
		return nil, err
	}

	// convert
	c := common.Identifier(cid)

	// return
	return &c, nil
}
