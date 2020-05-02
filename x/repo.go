package x

import (
	"encoding/json"
	"fmt"
	eth "github.com/ethereum/go-ethereum/ethclient"
	"github.com/floydeconomy/arisaedo-go/co"
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

// Add adds the case into IPFS and returns the identifier
func (r Repository) AddCase(c *Case) (*common.Identifier, error) {
	// error checks
	if err := c.Header().SanityCheck(); err != nil {
		return nil, err
	}
	if err := r.Verify(c.Header().CountryID()); err != nil {
		return nil, err
	}
	if err := r.Verify(c.Header().ProvinceID()); err != nil {
		return nil, err
	}

	// add to ipfs
	id, err := r.Put(c.Header().Body())
	if err != nil {
		return nil, err
	}

	// return
	return id, nil
}

// Verify verifies that a given identifier exists in IPFS
// todo: implement this (y error)?
func (r Repository) Verify(id common.Identifier) (y error) {
	err := r.shell.DagGet(string(id), &y)
	if err != nil {
		return
	}

	return
}

// todo: make this work wtih go routine
func (r Repository) VerifyCollection(c Collection) {
	for i, id := range c.Cases {
		// todo: remove println
		fmt.Printf("%v. verifying: [%s]\n", i, id)
		var goes co.Goes
		goes.Go(func() {
			r.Verify(id)
		})
	}
}