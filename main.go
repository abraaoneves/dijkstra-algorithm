package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"

	"com.github.abraaoneves/algorithm/dijkastra/graph"
)

func main() {
	input := CreateInput()
	gr := graph.CreateGraph(input)
	GetShortestPath(input.From, input.To, gr)
}

func CreateInput() graph.InputGraph {
	content, err := ioutil.ReadFile("./in.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var input graph.InputGraph
	err = json.Unmarshal(content, &input)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return input
}

func GetShortestPath(from, to string, gr *graph.Item) {
	startNode := &graph.Node{Value: from}
	endNode := &graph.Node{Value: to}

	path, distance := shortestPath(startNode, endNode, gr)

	fmt.Printf("Path: %v\n", path)
	fmt.Printf("Distance: %v\n", distance)
}

func shortestPath(start *graph.Node, end *graph.Node, item *graph.Item) ([]string, int) {
	visited := make(map[string]bool)
	distance := make(map[string]int)
	previous := make(map[string]string)

	queue := graph.Queue{}
	priorityQueue := queue.NewQueue()

	vertexStart := graph.Vertex{
		Node:     start,
		Distance: 0,
	}

	for _, value := range item.Nodes {
		distance[value.Value] = math.MaxInt64 / 2
	}

	distance[start.Value] = vertexStart.Distance
	priorityQueue.Enqueue(vertexStart)

	for !priorityQueue.IsEmpty() {
		value := priorityQueue.Dequeue()
		if visited[value.Node.Value] {
			continue
		}

		visited[value.Node.Value] = true
		near := item.Edges[*value.Node]

		for _, nearValue := range near {
			if !visited[nearValue.Node.Value] {
				if distance[value.Node.Value]+nearValue.Weight < distance[nearValue.Node.Value] {
					vertextStore := graph.Vertex{
						Node:     nearValue.Node,
						Distance: distance[value.Node.Value] + nearValue.Weight,
					}
					distance[nearValue.Node.Value] = distance[value.Node.Value] + nearValue.Weight
					previous[nearValue.Node.Value] = value.Node.Value
					priorityQueue.Enqueue(vertextStore)
				}
			}
		}
	}

	fmt.Printf("distance: %+v\n", distance)
	fmt.Printf("previous: %+v\n\n\n", previous)

	path := previous[end.Value]

	var arr []string
	arr = append(arr, end.Value)

	for path != start.Value {
		arr = append(arr, path)
		path = previous[path]
	}

	arr = append(arr, path)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr, distance[end.Value]
}
