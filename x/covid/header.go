package covid

import (
	"errors"
	"sync/atomic"
)

type Nonce [8]byte

// Header represents the entire COVID-19 case/death related data
type Header struct {
	// body
	body Body

	// cache
	cache struct {
		identifier atomic.Value
		country    atomic.Value
		province   atomic.Value
	}
}

// Body represents the cases/death related to COVID-19
// todo: implement signature and nonce fields
type Body struct {
	// IPFS Identifiers
	CountryID  Identifier `json:"country"`
	ProvinceID Identifier `json:"province"`

	// Case
	Time      uint64 `json:"time"`
	Confirmed uint64 `json:"confirmed"`
	Death     uint64 `json:"death"`
	Recovered uint64 `json:"recovered"`
	Active    uint64 `json:"active"`
}

func (h *Header) Confirmed() uint64      { return h.body.Confirmed }
func (h *Header) Death() uint64          { return h.body.Death }
func (h *Header) Recovered() uint64      { return h.body.Recovered }
func (h *Header) Active() uint64         { return h.body.Active }
func (h *Header) Time() uint64           { return h.body.Time }
func (h *Header) CountryID() Identifier  { return h.body.CountryID }
func (h *Header) ProvinceID() Identifier { return h.body.ProvinceID }
func (h *Header) Body() Body             { return h.body }

func (h *Header) SanityCheck() error {
	if h.Time() <= 0 {
		return errors.New("invalid time")
	}

	// ret: all passed
	return nil
}

// CopyBody creates a new header to mitigate any side-effects from
// modifying a body variable
func CopyBody(b *Body) *Body {
	cpy := *b
	return &cpy
}
