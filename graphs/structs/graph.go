package structs

type Graph[K comparable] struct {
	vertex             map[K]IVertex[K]
	edge               map[K]IEdge[K]
	vertexInstantiator IVertexInstantiator[K]
	edgesInstantiator  IEdgeInstantiator[K]
}

func NewGraph[K comparable](vertexInstantiator IVertexInstantiator[K], edgesInstantiator IEdgeInstantiator[K]) IGraph[K] {
	item := Graph[K]{
		vertex:             map[K]IVertex[K]{},
		edge:               map[K]IEdge[K]{},
		vertexInstantiator: vertexInstantiator,
		edgesInstantiator:  edgesInstantiator,
	}
	return &item
}

func (g Graph[K]) GetKeys() []K {
	return []K{}
}

func (g Graph[K]) GetVertex(key K) IVertex[K] {
	return g.vertex[key]
}

func (g Graph[K]) FindIncidentEdges(srcKey K, dstKey K) IEdge[K] {
	return g.edge[srcKey]
}

func (g *Graph[K]) AddEdge(srcKey K, dstKey K) IEdge[K] {
	srcVertex, isOk := g.vertex[srcKey]
	if !isOk || srcVertex == nil {
		srcVertex = g.vertexInstantiator.New(srcKey)
		g.vertex[srcKey] = srcVertex
	}

	dstVertex, isOk := g.vertex[dstKey]
	if !isOk || dstVertex == nil {
		dstVertex = g.vertexInstantiator.New(dstKey)
		g.vertex[dstKey] = dstVertex
	}

	e := g.edgesInstantiator.New(srcVertex, dstVertex)
	return e
}

func (g *Graph[K]) AddEdges(keys KeyCouples[K]) []IEdge[K] {
	result := []IEdge[K]{}
	for _, v := range keys {
		e := g.AddEdge(v.Src, v.Dst)
		result = append(result, e)
	}
	return result
}


