package main

import "fmt"

type Node struct {
	NodeID  string
	Peers   map[string]bool
	Storage map[string]bool
}

func NewNode(id string) *Node {
	return &Node{
		NodeID:  id,
		Peers:   make(map[string]bool),
		Storage: make(map[string]bool),
	}
}

func (n *Node) Broadcast(hash string) {
	for p := range n.Peers {
		fmt.Printf("Node %s broadcast to %s: %s\n", n.NodeID, p, hash)
	}
}
