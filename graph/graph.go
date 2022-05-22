package graph

import "sync"

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
	Edge  map[Node][]*Edge
	Lock  sync.RWMutex // TODO: Understand what is this <==
}
