package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Node struct {
	ID        int   `json:"id"`
	Neighbors []int `json:"neighbors"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
}

func main() {
	graph, err := loadGraph("graph.txt")
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}

	var startNode, endNode int
	fmt.Print("Enter start node: ")
	fmt.Scan(&startNode)
	fmt.Print("Enter end node: ")
	fmt.Scan(&endNode)

	fmt.Printf("Path from %d to %d:\n", startNode, endNode)
	path := findPath(graph, startNode, endNode)
	fmt.Println(path)
}

func findPath(graph Graph, startNode, endNode int) []int {
	visited := make(map[int]bool)
	var path []int
	dfsWithPath(graph, startNode, endNode, visited, &path)
	return path
}

func dfsWithPath(graph Graph, currentNode, endNode int, visited map[int]bool, path *[]int) {
	visited[currentNode] = true
	*path = append(*path, currentNode)

	if currentNode == endNode {
		return
	}

	for _, neighborID := range graph.Nodes[currentNode-1].Neighbors {
		if !visited[neighborID] {
			dfsWithPath(graph, neighborID, endNode, visited, path)
		}
	}

	if (*path)[len(*path)-1] != endNode {
		*path = (*path)[:len(*path)-1]
	}
}

func loadGraph(filename string) (Graph, error) {
	var graph Graph

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return graph, err
	}

	err = json.Unmarshal(fileContent, &graph)
	if err != nil {
		return graph, err
	}

	return graph, nil
}
