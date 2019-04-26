package hashring

import (
	"errors"
	"fmt"
	"hash/fnv"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/utils"
)

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

// HashRing is a consistency hash structure
type HashRing struct {
	replicate int

	nodeIds     *hashset.Set
	hashMapping *treemap.Map
	nodeMapping *treemap.Map
}

//Init initialize HashRing with physical node ids and replicate
func (hr *HashRing) Init(nodeIds []string, replicate int) error {
	if replicate <= 0 {
		return errors.New("replicate can not be zero")
	}
	hr.nodeIds = hashset.New()
	hr.replicate = replicate
	hr.hashMapping = treemap.NewWith(utils.UInt64Comparator)
	hr.nodeMapping = treemap.NewWithStringComparator()
	for _, node := range nodeIds {
		hr.AddNode(node)
	}
	return nil
}

// AddNode add a new server node to hash ring
func (hr *HashRing) AddNode(node string) {
	hr.nodeIds.Add(node)
	for i := 0; i < hr.replicate; i++ {
		vNode := fmt.Sprintf("%s#%d", node, i)
		hashCode := hash(vNode)
		hr.hashMapping.Put(hashCode, vNode)
		hr.nodeMapping.Put(vNode, node)
	}
}

// RemoveNode remove a server node from hash ring
func (hr *HashRing) RemoveNode(node string) {
	if hr.nodeIds.Contains(node) {
		hr.nodeIds.Remove(node)
		for i := 0; i < hr.replicate; i++ {
			vNode := fmt.Sprintf("%s#%d", node, i)
			hashCode := hash(vNode)
			hr.hashMapping.Remove(hashCode)
			hr.nodeMapping.Remove(vNode)
		}
	}
}

// SearchNodeForKey return server node witch store key
func (hr *HashRing) SearchNodeForKey(key uint64) string {
	it := hr.hashMapping.Iterator()
	for it.Begin(); it.Next(); {
		if key < it.Key().(uint64) {
			node, _ := hr.nodeMapping.Get(it.Value())
			return node.(string)
		}
	}
	it.First()
	node, _ := hr.nodeMapping.Get(it.Value())
	return node.(string)
}
