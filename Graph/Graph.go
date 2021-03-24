package Graph

import (
	"fmt"
)

// Node - graph node
type Node struct {
	Value int
	Links []*Node
}

// Graph data structure
type Graph struct {
	Nodes []*Node
}

// NewGraph - Graph constructor
func NewGraph() *Graph {
	nodes := make([]*Node, 0)
	return &Graph{nodes}
}

// NewNode - Node constructor
func NewNode(value int, links []*Node) *Node {
	return &Node{value, links}
}

// AddNode - add node to graph
func (g *Graph) AddNode(node *Node) {
	g.Nodes = append(g.Nodes, node)
}

// AddLink - Add link from one node to another
func (g *Graph) AddLink(start, end int) error {
	startNode, err := g.Find(start)
	if err != nil {
		return err
	}
	endNode, err := g.Find(end)
	if err != nil {
		return err
	}
	startNode.Links = append(startNode.Links, endNode)
	return nil
}

// Find node by it's value
func (g *Graph) Find(value int) (*Node, error) {
	for _, node := range g.Nodes {
		if node.Value == value {
			return node, nil
		}
	}
	return &Node{}, fmt.Errorf("Can't find graph node by value %d", value)
}
