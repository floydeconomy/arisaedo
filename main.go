package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//
// // Global variable to handle all the IPFS API client calls
// var sh *shell.Shell

func main() {
	// // Map of the Cases
	// s := store.StoreConfigure()
	//
	// // Infura
	// sh = ipfs.IPFSConfigure()
	//
	// // Inputs
	// c := c.CaseRandomAustralia()
	//
	// // Store
	// s.StoreCase(c)
	// fmt.Println(s)
	//
	// // Push to IPFS
	// s.StorePushPendingCases(sh)
	createServer()
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}

func createServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
