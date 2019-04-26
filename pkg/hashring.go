package hashring

import (
	"errors"
	"fmt"
	"hash/fnv"

	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/utils"
)

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

// HashRing is a consistency hash structure
type HashRing struct {
	nodeIds     []string
	replicate   int
	hashMapping *treemap.Map
	nodeMapping *treemap.Map
}

//Init initialize HashRing with physical node ids and replicate
func (hr *HashRing) Init(nodeIds []string, replicate int) error {
	if replicate <= 0 {
		return errors.New("replicate can not be zero")
	}

	hr.nodeIds = nodeIds
	hr.replicate = replicate
	hr.hashMapping = treemap.NewWith(utils.UInt64Comparator)
	hr.nodeMapping = treemap.NewWithStringComparator()
	for _, node := range nodeIds {
		for i := 0; i < hr.replicate; i++ {
			vNode := fmt.Sprintf("%d#%s", i, node)
			hashCode := hash(vNode)
			hr.hashMapping.Put(hashCode, vNode)
			hr.nodeMapping.Put(vNode, node)
		}
	}
	return nil
}
