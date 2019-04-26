package hashring

import (
	"testing"
)

func TestInit(t *testing.T) {
	hr := new(HashRing)
	err := hr.Init([]string{"1,1,1,1", "2,2,2,2"}, 3)
	if err != nil {
		t.Error(err)
	}
	t.Log("nodes:")
	for _, v := range hr.nodeIds {
		t.Log(v)
	}
	t.Log("virtual nodes:")
	it := hr.nodeMapping.Iterator()
	for it.Next() {
		t.Logf("%s -> %s", it.Key(), it.Value())
	}
	t.Log("hash codes:")
	it = hr.hashMapping.Iterator()
	for it.Next() {
		t.Logf("%d -> %s", it.Key(), it.Value())
	}
}
