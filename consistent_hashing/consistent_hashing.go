package consistent_hashing

import (
    "fmt"
    "hash/fnv"
)


type ConsistentHashingTable struct {
	hashRange uint32
	hashRing []NodeWrapper
}

type NodeWrapper struct {
	nodeStartBound uint32
	nodeInstance Node
}


// Constructor
func NewConsistentHashingTable(nodes []Node, config *ConsistentHashingConfig) *ConsistentHashingTable {

    defaultHashRangePerNode := config.getHashRange() / uint32(len(nodes))

    var hashRing []NodeWrapper
    for index, element := range nodes {
    	nodeWrapperInstance := NodeWrapper{uint32(index)*defaultHashRangePerNode, element}
	    hashRing = append(hashRing, nodeWrapperInstance)
    }

    fmt.Println(hashRing)

	table := &ConsistentHashingTable{config.getHashRange(), hashRing}
	return table
}


// Public API
func (table *ConsistentHashingTable) AddKey() {

}

func (table *ConsistentHashingTable) DeleteKey() {

}

func (table *ConsistentHashingTable) RetrieveKey(key string) interface{}{
	keyHash := hash(key) % table.hashRange
	serverInstance := table.binarySearchNode(keyHash, 0, uint32(len(table.hashRing) - 1))
	fmt.Println(keyHash)
	fmt.Println(serverInstance)
	return serverInstance.ProcessQueryByKey(key)
}

func (table *ConsistentHashingTable) AddNode() {
	
}


func (table *ConsistentHashingTable) DeleteNode(pos int) {
	if (pos >= len(table.hashRing) - 1) {
		return
	}
	tempHashRing := table.hashRing[:pos]

	table.hashRing = append(tempHashRing, table.hashRing[pos+1:]...)	
}

// Helper

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func (table *ConsistentHashingTable) binarySearchNode(hashValue uint32, lowerIndex uint32, upperIndex uint32) Node {

    if table.hashRing[upperIndex].nodeStartBound < hashValue {
    	return Node(table.hashRing[lowerIndex].nodeInstance)
    }

    if upperIndex - lowerIndex == 1 {
	    if table.hashRing[lowerIndex].nodeStartBound == hashValue {
	    	return Node(table.hashRing[lowerIndex].nodeInstance)
	    }
	    return Node(table.hashRing[upperIndex].nodeInstance)

    }

    mid := ((upperIndex - lowerIndex) / 2) + lowerIndex

    if table.hashRing[mid].nodeStartBound > hashValue {
    	return table.binarySearchNode(hashValue, lowerIndex, mid)
    } else if table.hashRing[mid].nodeStartBound < hashValue {
        return table.binarySearchNode(hashValue, mid, upperIndex)
    } else {
    	return Node(table.hashRing[mid].nodeInstance)
    }
}
