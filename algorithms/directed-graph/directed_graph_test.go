package directedgraph

import "testing"

func Test_Graph(t *testing.T) {
	AddEdge("1", "2")
	AddEdge("3", "4")
	ShowGraph()
}
