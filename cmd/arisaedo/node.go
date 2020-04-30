package arisaedo

import (
	"context"
	"encoding/json"
	eth "github.com/ethereum/go-ethereum/ethclient"
	"github.com/floydeconomy/arisaedo-go/pkg/co"
	"github.com/floydeconomy/arisaedo-go/pkg/common"
	"github.com/floydeconomy/arisaedo-go/x"
	shell "github.com/ipfs/go-ipfs-api"
	"log"
)

// todo: implement logDB
type Node struct {
	goes co.Goes

	shell *shell.Shell
	client *eth.Client
}

func New() *Node {
	s := shell.NewShell("https://ipfs.infura.io:5001")
	c, err := eth.Dial("https://mainnet.infura.io") // localhost:8545 for ganache-cli
	if err != nil {
		log.Fatal(err)
	}
	return &Node{
		shell:  s,
		client: c,
	}
}

func (n *Node) Run(ctx context.Context) error {
	// todo: implement houseKeeping and p2p comm.sync
	n.goes.Wait()
	return nil
}

// Put adds the interface to IPFS and returns the corresponding content identifier (CID)
func (n Node) Put(x interface {}) (*common.Identifier, error) {
	// marshall json
	m, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	// ipfs put
	cid, err := n.shell.DagPut(m, "json", "cbor")
	if err != nil {
		return nil, err
	}

	// convert
	c := common.Identifier(cid)

	// return
	return &c, nil
}

// Add adds the case into IPFS and returns the identifier
func (n Node) AddCase(c *x.Case) (*common.Identifier, error) {
	// error checks
	if err := c.Header().SanityCheck(); err != nil {
		return nil, err
	}
	if err := n.Verify(c.Header().CountryID()); err != nil {
		return nil, err
	}
	if err := n.Verify(c.Header().ProvinceID()); err != nil {
		return nil, err
	}

	// add to ipfs
	id, err := n.Put(c.Header().Body())
	if err != nil {
		return nil, err
	}

	// return
	return id, nil
}

// Verify verifies that a given identifier exists in IPFS
// todo: implement this
func (n Node) Verify(id common.Identifier) (y error) {
	err := n.shell.DagGet(string(id), &y)
	if err != nil {
		return
	}

	return
}