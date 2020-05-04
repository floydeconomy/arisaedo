package internal

import (
	"encoding/json"
	eth "github.com/ethereum/go-ethereum/ethclient"
	"github.com/floydeconomy/arisaedo-go/common"
	"github.com/floydeconomy/arisaedo-go/kv"
	shell "github.com/ipfs/go-ipfs-api"
	"log"
)

type Repository struct {
	data kv.Store

	shell *shell.Shell
	client *eth.Client
}

func New() *Repository {
	s := shell.NewShell("https://ipfs.infura.io:5001")
	c, err := eth.Dial("https://mainnet.infura.io") // localhost:8545 for ganache-cli
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{
		shell:  s,
		client: c,
	}
}

// Put adds the interface to IPFS and returns the corresponding content identifier (CID)
// todo: should Batch orders and put to kv store
func (r Repository) Put(x interface {}) (*common.Identifier, error) {
	// marshall json
	m, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	// ipfs put
	cid, err := r.shell.DagPut(m, "json", "cbor")
	if err != nil {
		return nil, err
	}

	// convert
	c := common.Identifier(cid)

	// return
	return &c, nil
}