package x

import (
	"github.com/ethereum/go-ethereum/common"
	c "github.com/floydeconomy/arisaedo-go/common"
	"sync/atomic"
)

// Collection aggregates COVID-19 cases
type Collection struct {
	ParentID  common.Hash `json:"parent identifier"`
	Timestamp uint64      `json:"timestamp"`

	Beneficiaries []common.Address // can be multiple
	Cases         c.Identifiers    // new cases

	cache struct {
		// todo: index should be strength of all previous collections
		index atomic.Value
	}
}

func (C *Collection) Verify() {
	// todo: implement verify for cases
}
