package structs

var _ IEdge[int] = SimpleEdge[int] {}

func NewSimpleEdge[K comparable](src IVertex[K], dst IVertex[K]) IEdge[K] {
	return &SimpleEdge[K]{
		src: src,
		dst: dst,
	}
}

type SimpleEdge[K comparable] struct {
	src IVertex[K]
	dst IVertex[K]
}

func (c SimpleEdge[K]) Src() IVertex[K] {
	return c.src
}

func (c SimpleEdge[K]) Dst() IVertex[K] {
	return c.dst
}

type SimpleEdgeInstantiator[K comparable] struct {

}

func (s SimpleEdgeInstantiator[K]) New(src IVertex[K], dst IVertex[K]) IEdge[K] {
	return NewSimpleEdge(src, dst)
}

