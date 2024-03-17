package main

import (
	"fmt"

	"github.com/aseerkt/go-dsa/pkg/graph"
)

func main() {
	// al := graph.AdjacencyList[string]{
	// 	"a": {"b", "c"},
	// 	"c": {"d"},
	// 	"d": {"e", "f"},
	// 	"f": {"h"},
	// 	"g": {"h", "i"},
	// 	"h": {"i"},
	// 	"i": {},
	// }

	// DEPTH FIRST PRINT
	// "a" -> "c" -> "d" -> "f" -> "h" -> "i"
	// "e" , "b"

	// al.DepthFirstPrint("a")

	// BREADTH FIRST PRINT
	// "a", "b", "c", "d", "e", "f", "h", "i"

	// al.BreadthFirstPrint("a")

	// HAS PATH

	// fmt.Println(al.HasPath("a", "d")) // true
	// fmt.Println(al.HasPath("a", "g")) // false

	// UNDIRECTED GRAPH

	el := graph.EdgeList[string]{
		{"a", "b"},
		{"b", "c"},
		{"c", "d"},
		{"e", "f"},
		{"f", "g"},
		{"g", "h"},
		{"h", "i"},
		{"j", "k"},
	}

	ual := el.GetAdjacentList()

	// CONNECTED COMPONENTS COUNT

	// fmt.Println("connected components count: ", ual.ConnectedComponentsCount())

	// LARGEST COMPONENT size

	fmt.Println("largest component size: ", ual.LargestComponentSize())
}
