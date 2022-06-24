package directedgraph

import "fmt"

var graph = make(map[string]map[string]bool)

func AddEdge(from string, to string) {
	edge := graph[from]
	if edge == nil {
		edge = make(map[string]bool)
		graph[from] = edge
	}
	graph[from][to] = true
}

func HasEdge(from string, to string) bool {
	return graph[from][to]
}

func ShowGraph() {
	fmt.Println(graph)
}
