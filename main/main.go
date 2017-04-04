package main

import (
    "fmt"
    "github.com/jlr52/exer-1-consistent-hashing/consistent_hashing"
    "github.com/jlr52/exer-1-consistent-hashing/server"
)

func main() {

	globalMap := map[string]string{
	    "key1": "key1",
	    "key2": "key2",
	    "key3": "key3",
	    "key4": "key4",
	    "key5": "key5"}

	server1 := server.NewFakeServer(1, globalMap)
	fmt.Println(server1)
	server2 := server.NewFakeServer(2, globalMap)
	//test, ok := server2.(consistent_hashing.Node)
	fmt.Println(server2)
	server3 := server.NewFakeServer(3, globalMap)
	//test, ok := server2.(consistent_hashing.Node)
	fmt.Println(server3)

	nodeList := []consistent_hashing.Node{server1, server2, server3}

	config := consistent_hashing.NewConsistentHashingConfig(uint32(100))

	table := consistent_hashing.NewConsistentHashingTable(nodeList, config)
	fmt.Println(table)
	fmt.Println(table.RetrieveKey("key1"))
    fmt.Println(table.RetrieveKey("key2"))
    fmt.Println(table.RetrieveKey("key3"))
    fmt.Println(table.RetrieveKey("ke4"))
    fmt.Println(table.RetrieveKey("key10"))
    table.DeleteNode(1)
    fmt.Println(table)
	fmt.Println(table.RetrieveKey("key1"))
    fmt.Println(table.RetrieveKey("key2"))
    fmt.Println(table.RetrieveKey("key3"))
    fmt.Println(table.RetrieveKey("ke4"))
    fmt.Println(table.RetrieveKey("key10"))
}