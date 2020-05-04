package common

type Identifier string

type Identifiers []Identifier

func IsEmpty(i Identifier) bool {
	return i == ""
}
