package x

import (
	"github.com/ethereum/go-ethereum/common"
	c "github.com/floydeconomy/arisaedo-go/common"
	"sync/atomic"
)

// Block aggregates identifiers that are signed by the beneficiaries
type Timeline struct {
	ParentID  common.Hash `json:"parent identifier"`
	Timestamp uint64      `json:"timestamp"`

	Beneficiaries []common.Address // can be multiple beneficiaries
	Identifiers   c.Identifiers    // new identifiers
	Update        []common.Hash	   // update wrong timeline-based data

	cache struct {
		index atomic.Value // beneficiaries identify index
	}
}