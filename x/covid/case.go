package covid

import (
	"sync/atomic"
)

// Identifier represents IPFS's CID
type Identifier string

// Case represents structure of COVID-19 cases
type Case struct {
	// body
	header *Header

	// caches
	cache struct {
		identifier atomic.Value
	}
}


//// NewCase creates a new cases
//func Compose(header *Header, body *Body) *Case {
//	c := &Case{header: CopyHeader(header)}
//	return c
//}

func (c *Case) Header() *Header { return c.header }
