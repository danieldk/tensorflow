package tensorflow

// #include "tensorflow.h"
import "C"

// Graph is a Tensorflow computation graph.
type Graph struct {
	graph *C.TF_Graph
}

// NewGraph allocates a graph.
func NewGraph() *Graph {
	return &Graph{
		graph: C.TF_NewGraph(),
	}
}

// Close deallocates a graph. Tensorflow automatically deallocates a graph
// when all sessions referencing the graph are closed.
func (g *Graph) Close() {
	C.TF_DeleteGraph(g.graph)
	g.graph = nil
}
