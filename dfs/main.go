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

	visited := make(map[int]bool)

	fmt.Println("DFS traversal:")
	for _, node := range graph.Nodes {
		if !visited[node.ID] {
			dfs(graph, node.ID, visited)
		}
	}
}

func dfs(graph Graph, currentNode int, visited map[int]bool) {
	visited[currentNode] = true
	fmt.Println(currentNode)

	for _, neighborID := range graph.Nodes[currentNode-1].Neighbors {
		if !visited[neighborID] {
			dfs(graph, neighborID, visited)
		}
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
