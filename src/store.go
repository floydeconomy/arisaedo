package store

import (
	"encoding/json"
	"fmt"
)

type Store struct {
	Pending map[string]Case        `json:"Pending"`
	Address map[string]Addressable `json:"Address"`
}

func StoreConfigure() *Store {
	return &Store{
		Pending: make(map[string]Case),
		Address: make(map[string]Addressable),
	}
}

func (s *Store) StorePushPendingCases() {
	// Marshall JSON
	m, err := json.Marshal(s.Pending)
	if err != nil {
		fmt.Println(err)
	}

	// // Push to IPFS
	// start := time.Now()
	// cid, err := sh.DagPut(m, "json", "cbor")
	// elapsed := time.Since(start)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("WRITE: IPLD PUT call took ", elapsed)
	//
	// // Make Address
	// s.Address["12345"] = Addressable{
	//   CIDs: make(map[string]CID)
	// }
	//
	// s.Address["12345"].CIDs[cid] = CID{
	//   Timestamp = t.Time.Now()
	// }
}
