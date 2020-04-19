package main

import (
	"fmt"

	c "github.com/floydeconomy/theCovidInitiative/x/cases"
	ipfs "github.com/floydeconomy/theCovidInitiative/x/ipfs"
	store "github.com/floydeconomy/theCovidInitiative/x/store"
	shell "github.com/ipfs/go-ipfs-api"
)

// Global variable to handle all the IPFS API client calls
var sh *shell.Shell

func main() {
	// Map of the Cases
	s := store.StoreConfigure()

	// Infura
	sh = ipfs.IPFSConfigure()

	// Inputs
	c := c.CaseRandomAustralia()

	// Store
	s.StoreCase(c)
	fmt.Println(s)

	// Push to IPFS
	s.StorePushPendingCases(sh)
}
