package store

//
//// Getter defines the method to read the data store
//type Getter interface {
//	Get(key []byte) ([]byte, error)
//	Has(key []byte) (bool, error)
//}
//
//// Putter defines the method to put into data store
//type Putter interface {
//	Put(value []byte) (string, error)
//	Delete(key []byte) error
//}
//
//// Pair defines the key-value persistent pair
//type Pair interface {
//	Key() []byte
//	Value() []byte
//}

// Operator defines the store chain interface
type DataOperator interface {
	IDataStore
	Close() error
}

// Store defines the functional implementation of a legitimate key-value store
type IDataStore interface {
	Getter
	Putter
}
