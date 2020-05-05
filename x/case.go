package x

import (
	"errors"
	"github.com/floydeconomy/arisaedo-go/common"
)

type Cases []Case

// Case represents structure of COVID-19 cases
type Case struct {
	Header *CaseHeader

	cache struct {
		identifier string
	}
}

// CaseHeader represent store information about the COVID-19 cases
type CaseHeader struct {
	// identifiers
	CountryID  common.Identifier `json:"country"`

	// body
	Time      uint64 `json:"time"`
	Confirmed uint64 `json:"confirmed"`
	Death     uint64 `json:"death"`
	Recovered uint64 `json:"recovered"`
	Active    uint64 `json:"active"`
}

// Compose method is usually to recover a case by its portions
func Compose(header *CaseHeader) *Case {
	return &Case{
		Header: header,
	}
}

// SanityCheck checks whether the case is valid
func (c *Case) SanityCheck() error {
	if c.Header.Time == 0 {
		return errors.New("invalid time")
	}
	if common.IsEmpty(c.Header.CountryID) {
		return errors.New("country id doesn't exists")
	}
	return nil
}

// Identifier represent the IPFS identifier for this case
func (c *Case) Identifier() string {
	// todo: ensure this is never returning empty string or undefined
	return c.cache.identifier
}

//// Add adds the case into IPFS and returns the identifier
//func AddCase(c *Case, r *Repository) (*common.Identifier, error) {
//	// error checks
//	if err := c.Header().SanityCheck(); err != nil {
//		return nil, err
//	}
//	if err := r.Verify(c.Header().CountryID()); err != nil {
//		return nil, err
//	}
//	if err := r.Verify(c.Header().ProvinceID()); err != nil {
//		return nil, err
//	}
//
//	// add to ipfs
//	id, err := r.Put(c.Header().Body())
//	if err != nil {
//		return nil, err
//	}
//
//	// return
//	return id, nil
//}