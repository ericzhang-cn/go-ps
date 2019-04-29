package kvstore

import (
	"os"
	"testing"
)

func TestGetAndPut(t *testing.T) {
	kv := new(BadgerStore)
	kv.Dir = "./data"
	err := kv.Put(uint64(42), []byte("hello world"))
	if err != nil {
		t.Error(err)
	}
	value, err := kv.Get(uint64(42))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", value)
	os.RemoveAll("./data")
}

func TestGetBatchAndBatchPutBatch(t *testing.T) {
	kv := new(BadgerStore)
	kv.Dir = "./data"
	keys := []uint64{1, 2, 3}
	values := [][]byte{[]byte("Alice"), []byte("Bob"), []byte("Oscar")}
	err := kv.PutBatch(keys, values)
	if err != nil {
		t.Error(err)
	}
	v, err := kv.GetBatch(keys)
	if err != nil {
		t.Error(err)
	}
	for _, value := range v {
		t.Logf("%s", value)
	}
	os.RemoveAll("./data")
}

func TestGetRange(t *testing.T) {
	kv := new(BadgerStore)
	kv.Dir = "./data"
	keys := []uint64{10, 15, 16}
	values := [][]byte{[]byte("Alice"), []byte("Bob"), []byte("Oscar")}
	err := kv.PutBatch(keys, values)
	if err != nil {
		t.Error(err)
	}
	v, err := kv.GetRange(10, 20)
	if err != nil {
		t.Error(err)
	}
	for _, value := range v {
		t.Logf("%s", value)
	}
	os.RemoveAll("./data")
}
