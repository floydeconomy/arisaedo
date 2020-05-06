package kv

import (
	"github.com/floydeconomy/arisaedo-go/store/data"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	writeOpt = opt.WriteOptions{}
	readOpt  = opt.ReadOptions{}
	scanOpt  = opt.ReadOptions{DontFillCache: true}
)

type lvldb struct {
	db *leveldb.DB
}

func New(db *leveldb.DB) data.Operator {
	return &lvldb{db}
}

func (ldb *lvldb) Close() error {
	return ldb.db.Close()
}

func (ldb *lvldb) IsNotFound(err error) bool {
	return err == leveldb.ErrNotFound
}

func (ldb *lvldb) Get(key []byte) ([]byte, error) {
	val, err := ldb.db.Get(key, &readOpt)
	// val will be []byte{} if error occurs, which is not expected
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (ldb *lvldb) Has(key []byte) (bool, error) {
	return ldb.db.Has(key, &readOpt)
}

func (ldb *lvldb) Put(key, val []byte) error {
	return ldb.db.Put(key, val, &writeOpt)
}

func (ldb *lvldb) Delete(key []byte) error {
	return ldb.db.Delete(key, &writeOpt)
}
