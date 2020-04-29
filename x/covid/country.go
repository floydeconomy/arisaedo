package covid

// Country represents locations of countries affected by COVID-19
type Country struct {
	Name     string     `json:"Name"`
	Location Coordinate `json:"Location"`
}

// Coordinate represents lat/lng coordinates of a location
type Coordinate struct {
	Latitude  int8 `json:"Latitude"`
	Longitude int8 `json:"Longitude"`
}

