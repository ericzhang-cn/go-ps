package server

import (
	"github.com/dgraph-io/badger"
)

// BadgerStore is kv store implementation using [badger](https://github.com/dgraph-io/badger)
type BadgerStore struct {
	dir string
}

func (bs *BadgerStore) open() (db *badger.DB, err error) {
	opts := badger.DefaultOptions
	opts.Dir = bs.dir
	opts.ValueDir = bs.dir
	db, err = badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Get retrieves value of the key
func (bs *BadgerStore) Get(key uint64) (value []byte, err error) {
	return nil, nil
}

// Put add or updates value of key
func (bs *BadgerStore) Put(key uint64, value []byte) (err error) {
	return nil
}

// GetBatch retrieves a list of values of keys
func (bs *BadgerStore) GetBatch(keys []uint64) (values [][]byte, err error) {
	return nil, nil
}

// PutBatch add or updates a list of key-value pairs
func (bs *BadgerStore) PutBatch(keys []uint64, values [][]byte) (err error) {
	return nil
}
