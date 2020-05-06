package chain

// define individual functions for data store

type (
	GetFunc    func(key []byte) ([]byte, error)
	HasFunc    func(key []byte) (bool, error)
	PutFunc    func(key, value []byte) error
	DeleteFunc func(key []byte) error
	KeyFunc    func() []byte
	ValueFunc  func() []byte
)

func (f GetFunc) Get(key []byte) ([]byte, error) { return f(key) }
func (f HasFunc) Has(key []byte) (bool, error)   { return f(key) }
func (f PutFunc) Put(key, value []byte) error    { return f(key, value) }
func (f DeleteFunc) Delete(key []byte) error     { return f(key) }
func (f KeyFunc) Key() []byte                    { return f() }
func (f ValueFunc) Value() []byte                { return f() }
