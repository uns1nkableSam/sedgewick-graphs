package structs

var _ IVertex[int] = SimpleVertex[int] {}

func NewSimpleVertex[K comparable](key K) IVertex[K] {
	return &SimpleVertex[K]{
		key: key,
		edges: []IEdge[K]{},
	}
}

type SimpleVertex[K comparable] struct {
	key K
	edges []IEdge[K]
}

func (c SimpleVertex[K]) GetKey() K {
	return c.key
}

func (c SimpleVertex[K]) GetEdges() []IEdge[K] {
	return c.edges
}

func (c SimpleVertex[K]) GetDegree() int {
	return len(c.edges)
}

type SimpleVertexInstantiator[K comparable] struct {

}

func (s SimpleVertexInstantiator[K]) New(key K) IVertex[K] {
	return NewSimpleVertex[K](key)
}