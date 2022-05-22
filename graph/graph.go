package graph

import (
	"sync"
)

type InputData struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
}

type InputGraph struct {
	Graph []InputData `json:"graph"`
	From  string      `json:"from"`
	To    string      `json:"to"`
}

type Item struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	Lock  sync.RWMutex
}

func (item *Item) AddNode(node *Node) {
	item.Lock.Lock()
	item.Nodes = append(item.Nodes, node)
	item.Lock.Unlock()
}

func (item *Item) AddEdge(nodeOne, nodeTwo *Node, weight int) {
	item.Lock.Lock()

	if item.Edges == nil {
		item.Edges = make(map[Node][]*Edge)
	}

	edgeOne := Edge{
		Node:   nodeTwo,
		Weight: weight,
	}

	edgeTwo := Edge{
		Node:   nodeOne,
		Weight: weight,
	}

	item.Edges[*nodeOne] = append(item.Edges[*nodeOne], &edgeOne)
	item.Edges[*nodeTwo] = append(item.Edges[*nodeTwo], &edgeTwo)

	item.Lock.Unlock()
}
