package graph

type EdgeList[V comparable] [][2]V

func (el *EdgeList[V]) GetAdjacentList() *AdjacencyList[V] {
	al := AdjacencyList[V]{}
	for _, pairs := range *el {
		al[pairs[0]] = append(al[pairs[0]], pairs[1])
		al[pairs[1]] = append(al[pairs[1]], pairs[0])
	}
	return &al
}
