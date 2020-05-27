package store

type Store struct {
	db     ChainOperator
	client DataOperator
}

type Options struct {
	Db    string
	Chain string
}

func New(o Options) *Store {
	return &Store{
		db:     newDataStore(o.Db),
		client: newChainStore(o.Chain),
	}
}
