package server

import (
    "strconv"
    //"fmt"
)

type FakeServer struct {
	id int
	store map[string]string
}


func NewFakeServer(id int, globalMap map[string]string) *FakeServer{

    localCopyGlobalMap := make(map[string]string)

    for k,v := range globalMap {
    	localCopyGlobalMap[k] = v
    }

	server := &FakeServer{id, localCopyGlobalMap}

    return server
} 


func (server *FakeServer) ProcessQueryByKey(key string) interface{} {
	return map[string]string{
	    "id": strconv.Itoa(server.id),
	    "value":   server.store[key]}
}
