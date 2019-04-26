package hashring

import (
	"testing"
)

func Test(t *testing.T) {
	hr := new(HashRing)
	err := hr.Init([]string{"1.1.1.1", "2.2.2.2"}, 16)
	if err != nil {
		t.Error(err)
	}
	t.Log("virtual nodes:")
	it := hr.nodeMapping.Iterator()
	for it.Begin(); it.Next(); {
		t.Logf("%s -> %s", it.Key(), it.Value())
	}
	t.Log("hash codes:")
	it = hr.hashMapping.Iterator()
	for it.Begin(); it.Next(); {
		t.Logf("%d -> %s", it.Key(), it.Value())
	}

	hr.AddNode("4.4.4.4")
	hr.RemoveNode("1.1.1.1")
	t.Log("virtual nodes:")
	it = hr.nodeMapping.Iterator()
	for it.Begin(); it.Next(); {
		t.Logf("%s -> %s", it.Key(), it.Value())
	}
	t.Log("hash codes:")
	it = hr.hashMapping.Iterator()
	for it.Begin(); it.Next(); {
		t.Logf("%d -> %s", it.Key(), it.Value())
	}

	t.Logf("key: %d, server node: %s", uint64(0), hr.SearchNodeForKey(uint64(0)))
	t.Logf("key: %d, server node: %s", uint64(14854670415487301923), hr.SearchNodeForKey(uint64(14854670415487301923)))
	t.Logf("key: %d, server node: %s", uint64(14854670415487301924), hr.SearchNodeForKey(uint64(14854670415487301924)))
	t.Logf("key: %d, server node: %s", uint64(14854673714022186558), hr.SearchNodeForKey(uint64(14854673714022186558)))
}
