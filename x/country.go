package x

import "sync/atomic"

// Country represents locations of countries affected by COVID-19
type Country struct {
	header *LocationHeader

	// caches
	identifier atomic.Value
}

// Province represents states in a country
type Province struct {
	header *LocationHeader

	// caches
	identifier atomic.Value
}

// LocationHeader represents vital information about a certain place
type LocationHeader struct {
	Name     string     `json:"Name"`
	Location Coordinate `json:"Location"`
}

// Coordinate represents lat/lng coordinates of a location
type Coordinate struct {
	Latitude  int8 `json:"Latitude"`
	Longitude int8 `json:"Longitude"`
}

