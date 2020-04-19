package store

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	c "github.com/floydeconomy/theCovidInitiative/x/cases"
	shell "github.com/ipfs/go-ipfs-api"
)

type Store struct {
	Pending map[string]c.Case        `json:"Pending"`
	Address map[string]c.Addressable `json:"Address"`
}

func StoreConfigure() *Store {
	return &Store{
		Pending: make(map[string]c.Case),
		Address: make(map[string]c.Addressable),
	}
}

func (s *Store) StorePushPendingCases(sh *shell.Shell) {
	// Marshall JSON
	m, err := json.Marshal(s.Pending)
	if err != nil {
		fmt.Println(err)
	}

	// Push to IPFS
	start := time.Now()
	cid, err := sh.DagPut(m, "json", "cbor")
	elapsed := time.Since(start)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println("WRITE: IPLD PUT call took ", elapsed)
	fmt.Println("CID: ", cid)
	// // Make Address
	// s.Address["12345"] = Addressable{
	//   CIDs: make(map[string]CID)
	// }
	//
	// s.Address["12345"].CIDs[cid] = CID{
	//   Timestamp = t.Time.Now()
	// }
}

func (s *Store) StoreCase(c c.Case) {
	s.Pending[c.Country] = c
}
