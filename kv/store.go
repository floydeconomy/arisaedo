package kv

// Getter defines the method to read the kv store
type Getter interface {
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
}

// Putter defines the method to put into kv store
type Putter interface {
	Put(key, value []byte) error
	Delete(key []byte) error
}

// Pair defines the key-value persistent pair
type Pair interface {
	Key() []byte
	Value() []byte
}

// Operator defines the store operator interface
type Operator interface {
	Store
	Close() error
}

// Store defines the functional implementation of a legitatimate key-value store
type Store interface {
	Getter
	Putter
}
