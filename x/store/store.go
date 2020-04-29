package store

import (
	"encoding/json"
	"github.com/floydeconomy/arisaedo-go/x/covid"
	shell "github.com/ipfs/go-ipfs-api"
)

// Store represents the main connections to ipfs and ethereum blockchain
type Store struct {
	Shell *shell.Shell
}

// Compose makes a store structure which maintains connections
// to IPFS through the shell variable
func Compose() *Store {
	return &Store{
		Shell: shell.NewShell("https://ipfs.infura.io:5001"),
	}
}

// Put adds the interface to IPFS and returns the corresponding content identifier (CID)
func (s Store) Put(x interface {}) (*covid.Identifier, error) {
	// marshall json
	m, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	// ipfs put
	cid, err := s.Shell.DagPut(m, "json", "cbor")
	if err != nil {
		return nil, err
	}

	// convert
	c := covid.Identifier(cid)

	// return
	return &c, nil
}

// Add adds the case into IPFS and returns the identifier
func (s Store) AddCase(c *covid.Case) (*covid.Identifier, error) {
	// checks
	err := c.Header().SanityCheck()
	if err != nil {
		return nil, err
	}

	// verify identifiers
	err = s.Verify(c.Header().CountryID())
	if err != nil {
		return nil, err
	}

	err = s.Verify(c.Header().ProvinceID())
	if err != nil {
		return nil, err
	}

	// add to ipfs
	id, err := s.Put(c.Header().Body())
	if err != nil {
		return nil, err
	}

	// return
	return id, nil
}

// Verify verifies that a given identifier exists in IPFS
// todo: implement this
func (s Store) Verify(id covid.Identifier) error {
	return nil
}