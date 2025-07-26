package main

import "sort"

const murmurSeed uint32 = 42

type HashRing struct {
	Nodes []*ServerNode
}

func NewHashRing(nodes []*ServerNode) *HashRing {
	copied := make([]*ServerNode, len(nodes))
	copy(copied, nodes)
	return &HashRing{Nodes: copied}
}

func (hr *HashRing) AddNode(node *ServerNode) {
	hr.Nodes = append(hr.Nodes, node)
	hr.sortNodes()
}

func (hr *HashRing) RemoveNode(nodeID string) {
	for i, node := range hr.Nodes {
		if node.ID == murmur3_32([]byte(nodeID), murmurSeed) {
			hr.Nodes = append(hr.Nodes[:i], hr.Nodes[i+1:]...)
			break
		}
	}
	hr.sortNodes()
}

func (hr *HashRing) AssignKeyToNode(key string) *ServerNode {
	node := hr.GetNearestNode(key)

	node.Keys = append(node.Keys, key)
	return node
}

func (hr *HashRing) GetNearestNode(key string) *ServerNode {
	keyHash := murmur3_32([]byte(key), murmurSeed)
	for i := range hr.Nodes {
		if keyHash <= hr.Nodes[i].ID {
			return hr.Nodes[i]
		}
	}
	return hr.Nodes[0]
}

func (hr *HashRing) sortNodes() {
	sort.Slice(hr.Nodes, func(i, j int) bool {
		return hr.Nodes[i].ID < hr.Nodes[j].ID
	})
}

func (hr *HashRing) FindNodeByID(nodeID string) *ServerNode {
	id := murmur3_32([]byte(nodeID), murmurSeed)
	for i := range hr.Nodes {
		if hr.Nodes[i].ID == id {
			return hr.Nodes[i]
		}
	}
	return nil
}
