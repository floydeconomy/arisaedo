package main

import "fmt"

func PrintAPIMessage(apiURL string, nodeID string) {
	fmt.Printf(`    API portal   [ %v ]
    Node ID      [ %v ]
`,
		apiURL,
		nodeID)
}

func PrintStoreMessage(ipfs string, eth string) {
	fmt.Printf(`    IPFS Endpoint   [ %v ]
    Ethereum Client      [ %v ]
`,
		ipfs,
		eth)
}
