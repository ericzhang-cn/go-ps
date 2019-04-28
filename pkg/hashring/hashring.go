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

// HashRing is internal data structure of consistency hash ring
type HashRing struct {
	numOfReplicate uint32

	nodeIdentities        *hashset.Set
	hashCodeToVirtualNode *treemap.Map
	virtualNodeToNode     *treemap.Map
}

//Init initializes HashRing with number of replicate
func (hr *HashRing) Init(numOfReplicate uint32) error {
	if numOfReplicate == 0 {
		return errors.New("replicate can not be zero")
	}
	hr.nodeIdentities = hashset.New()
	hr.numOfReplicate = numOfReplicate
	hr.hashCodeToVirtualNode = treemap.NewWith(utils.UInt64Comparator)
	hr.virtualNodeToNode = treemap.NewWith(utils.StringComparator)
	return nil
}

// AddNode add a new server node to hash ring
func (hr *HashRing) AddNode(node string) error {
	if hr.nodeIdentities.Contains(node) {
		return errors.New("Server node already exists")
	}
	hr.nodeIdentities.Add(node)
	for i := uint32(0); i < hr.numOfReplicate; i++ {
		vNode := fmt.Sprintf("%s#%d", node, i)
		hashCode := hash(vNode)
		hr.hashCodeToVirtualNode.Put(hashCode, vNode)
		hr.virtualNodeToNode.Put(vNode, node)
	}
	return nil
}

// RemoveNode removes a server node from hash ring
func (hr *HashRing) RemoveNode(node string) error {
	if !hr.nodeIdentities.Contains(node) {
		return errors.New("Server node not found")
	}
	if hr.nodeIdentities.Contains(node) {
		hr.nodeIdentities.Remove(node)
		for i := uint32(0); i < hr.numOfReplicate; i++ {
			vNode := fmt.Sprintf("%s#%d", node, i)
			hashCode := hash(vNode)
			hr.hashCodeToVirtualNode.Remove(hashCode)
			hr.virtualNodeToNode.Remove(vNode)
		}
	}
	return nil
}

// SearchForKey return server node that stores  key
func (hr *HashRing) SearchForKey(key uint64) (string, error) {
	if hr.nodeIdentities.Size() == 0 {
		return "", errors.New("Empty hash ring")
	}
	it := hr.hashCodeToVirtualNode.Iterator()
	for it.Begin(); it.Next(); {
		if key < it.Key().(uint64) {
			node, _ := hr.virtualNodeToNode.Get(it.Value())
			return node.(string), nil
		}
	}
	it.First()
	node, _ := hr.virtualNodeToNode.Get(it.Value())
	return node.(string), nil
}
