package kvstore

import (
	"encoding/binary"
	"errors"

	"github.com/dgraph-io/badger"
)

// BadgerStore is kv store implementation using badger
type BadgerStore struct {
	Dir string
}

func (bs *BadgerStore) open() (db *badger.DB, err error) {
	opts := badger.DefaultOptions
	opts.Dir = bs.Dir
	opts.ValueDir = bs.Dir
	db, err = badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Get retrieves value of the key
func (bs *BadgerStore) Get(key uint64) (value []byte, err error) {
	db, err := bs.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		buf := make([]byte, 8)
		binary.PutUvarint(buf, key)
		item, err := txn.Get(buf)
		if err != nil {
			return err
		}
		value, err = item.Value()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Put add or updates value of key
func (bs *BadgerStore) Put(key uint64, value []byte) (err error) {
	db, err := bs.open()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		buf := make([]byte, 8)
		binary.PutUvarint(buf, key)
		err := txn.Set(buf, value)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

// GetBatch retrieves a list of values of keys
func (bs *BadgerStore) GetBatch(keys []uint64) (values [][]byte, err error) {
	db, err := bs.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		values = make([][]byte, len(keys))
		for i, key := range keys {
			buf := make([]byte, 8)
			binary.PutUvarint(buf, key)
			item, err := txn.Get(buf)
			if err != nil {
				return err
			}
			values[i], err = item.ValueCopy(nil)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

// PutBatch add or updates a list of key-value pairs
func (bs *BadgerStore) PutBatch(keys []uint64, values [][]byte) (err error) {
	db, err := bs.open()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		for i, key := range keys {
			buf := make([]byte, 8)
			binary.PutUvarint(buf, key)
			err := txn.Set(buf, values[i])
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// GetRange retrieves data of key range from begin (included) to end (excluded)
func (bs *BadgerStore) GetRange(begin uint64, end uint64) (values [][]byte, err error) {
	if begin >= end {
		return nil, errors.New("begin have to be less than end")
	}
	db, err := bs.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		values = make([][]byte, end-begin)
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()
		buf := make([]byte, 8)
		binary.PutUvarint(buf, begin)
		for it.Seek(buf); it.Valid(); it.Next() {
			item := it.Item()
			k, _ := binary.Uvarint(item.Key())
			if k >= end {
				break
			}
			values[k-begin], err = item.Value()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return values, nil
}
