package main

import "fmt"

func main() {
	node1 := ServerNode{ID: murmur3_32([]byte("node1"), 42), Addr: "localhost:8080"}
	node2 := ServerNode{ID: murmur3_32([]byte("node2"), 42), Addr: "localhost:8081"}
	node3 := ServerNode{ID: murmur3_32([]byte("node3"), 42), Addr: "localhost:8082"}
	hashRing := NewHashRing([]*ServerNode{&node1, &node2, &node3})

	request := "asdkasdkj@a"

	n := hashRing.AssignKeyToNode(string(request))
	fmt.Printf("Assigned node: %s with ID: %d\n", n.Addr, n.ID)

}
