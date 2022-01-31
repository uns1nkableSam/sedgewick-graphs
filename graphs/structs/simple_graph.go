package structs

func NewSimpleGraph[K comparable]() IGraph[K] {
	vertices := &SimpleVertexInstantiator[K]{}
	edges := &SimpleEdgeInstantiator[K]{}

	item := NewGraph[K](vertices, edges)
	return item
}
