package cases

import (
	c "github.com/ethereum/go-ethereum/common"
	"github.com/floydeconomy/theCovidInitiative/pkg/common"
)

type Person struct {
	Address c.Address
	Links   Addressable
}

type Addressable struct {
	Identifiers map[string]common.Identifier `json:"Content Identifiers"`
}
