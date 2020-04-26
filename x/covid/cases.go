package cases

import (
	t "time"

	"github.com/floydeconomy/theCovidInitiative/pkg/common"
)

type Case struct {
	ID        common.Identifier `json:"CID"`
	Country   string            `json:"Country"`
	Total     string            `json:"ID"`
	Active    string            `json:"Name"`
	Death     string            `json:"Salary"`
	Timestamp t.Time            `json:"Time"`
}
