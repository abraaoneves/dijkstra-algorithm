package graph

import "sync"

type PriorityQueue []*Vertex

type Queue struct {
	Items []Vertex
	Lock  sync.RWMutex
}

const FIRST int = 0

// Enqueue adds an Node to the end of the queue
func (queue *Queue) Enqueue(vertex Vertex) {
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

// Remove an item from starts of the queue
func (queue *Queue) Dequeue() *Vertex {
	queue.Lock.Lock()
	item := queue.Items[FIRST]
	queue.Items = queue.Items[1:len(queue.Items)]
	queue.Lock.Unlock()
	return &item
}

// Creates new Queue
func (queue *Queue) NewQueue() *Queue {
	queue.Lock.Lock()
	queue.Items = []Vertex{}
	queue.Lock.Unlock()
	return queue
}

func (queue *Queue) isEmpty() bool {
	queue.Lock.RLock()
	defer queue.Lock.RUnlock()
	return len(queue.Items) == 0
}

func (queue *Queue) Size() int {
	queue.Lock.RLock()
	defer queue.Lock.RUnlock()
	return len(queue.Items)
}
