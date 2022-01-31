package main

import (
	"fmt"

	"uns1nkable.sam/sedgewick/graphs/graphs/structs"
)

func main() {
	graph := structs.NewSimpleGraph[int]()
	graph.AddEdge(1, 2)
	graph.AddEdges(structs.KeyCouples[int]{{1, 3},{2, 4}})
	fmt.Printf("%v", graph)
}
