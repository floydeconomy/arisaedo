package covid

import "sync/atomic"

// Identifier represents IPFS's CID
type Identifier string

// Case represents structure of COVID-19 cases
type Case struct {
	// header
	header *Header

	// body
	body *Body

	// caches
	identifier atomic.Value
}

// Header represents the vital information associated with the COVID-19 case
type Header struct {
	Country Country `json:"country"`
}

// Body represents the cases/death related to COVID-19
type Body struct {
	Confirmed uint64  `json:"confirmed"`
	Death     uint64  `json:"death"`
	Recovered uint64  `json:"recovered"`
	Active    uint64  `json:"active"`
	Time      uint64  `json:"time"`
	Nonce     [8]byte `json:"nonce"` // todo: how does this change?
}
