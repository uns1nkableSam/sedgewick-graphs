package structs

type KeyCouple[K comparable] struct {
	Src, Dst K
}

type KeyCouples[K comparable] []KeyCouple[K]

type IVertexInstantiator[K comparable] interface {
	New(key K) IVertex[K]
}

type IEdgeInstantiator[K comparable] interface {
	New(src IVertex[K], dst IVertex[K]) IEdge[K]
}

type IGraph [K comparable] interface {
	GetKeys()[]K
	GetVertex(key K)IVertex[K]
	FindIncidentEdges(srcKey K, dstKey K) IEdge[K]
	AddEdge(src K, dst K) IEdge[K]
	AddEdges(keys KeyCouples[K]) []IEdge[K]
}

type IVertex[K comparable] interface {
	GetKey() K
	GetEdges() []IEdge[K]
	GetDegree() int
}

type IEdge[K comparable] interface {
	Src() IVertex[K]
	Dst() IVertex[K]
}

type IDataContainer [T any] interface {
	GetData() *T
}

func IsDataContainerOf[T any] (i interface{}) (IDataContainer[T], bool) {
	cont, isOk := i.(IDataContainer[T])
	if !isOk || cont == nil {
		return nil, isOk
	}
	return cont, true
}
