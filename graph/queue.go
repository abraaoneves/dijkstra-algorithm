package graph

import "sync"

type PriorityQueue []*Vertex

type NodeQueue struct {
	Items []Vertex
	Lock  sync.RWMutex
}

// Enqueue adds an Node to the end of the queue
func (queue *NodeQueue) Enqueue(vertex Vertex) {
	queue.Lock.Lock()

	// add first value to the queue
	if len(queue.Items) == 0 {
		queue.Items = append(queue.Items, vertex)
		queue.Lock.Unlock()
		return
	}

	var isInserted bool
	for key, value := range queue.Items {
		if vertex.Distance < value.Distance {
			if key > 0 {
				queue.Items = append(queue.Items[:key+1], queue.Items[key:]...)
				queue.Items[key] = vertex
				isInserted = true
			} else {
				queue.Items = append([]Vertex{vertex}, queue.Items...)
				isInserted = true
			}
		}

		if isInserted {
			break
		}
	}

	if !isInserted {
		queue.Items = append(queue.Items, vertex)
	}

	queue.Lock.Unlock()
}
