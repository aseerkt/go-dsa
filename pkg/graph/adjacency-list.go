package graph

import (
	"fmt"

	"github.com/aseerkt/go-dsa/pkg/queue"
	"github.com/aseerkt/go-dsa/pkg/stack"
)

type AdjacencyList[V comparable] map[V][]V

func (al *AdjacencyList[V]) DepthFirstPrint(src V) {
	st := stack.New[V]()
	st.Push(src)

	visited := make(map[V]bool)

	visited[src] = true

	for !st.IsEmpty() {
		current := st.Pop()

		fmt.Println(current)

		neighbors := (*al)[current]

		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				st.Push(neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func (al *AdjacencyList[V]) BreadthFirstPrint(src V) {
	qu := queue.New[V]()
	qu.Enqueue(src)

	visited := make(map[V]bool)

	for !qu.IsEmpty() {
		current := qu.Dequeue()
		visited[current] = true

		fmt.Println(current)

		neighbors := (*al)[current]

		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				qu.Enqueue(neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func (al *AdjacencyList[V]) HasPath(src V, dst V) bool {
	qu := queue.New[V]()
	qu.Enqueue(src)

	visited := make(map[V]bool)

	for !qu.IsEmpty() {
		current := qu.Dequeue()
		visited[current] = true

		fmt.Println("current, dst", current, dst)

		if current == dst {
			return true
		}

		neighbors := (*al)[current]

		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				qu.Enqueue(neighbor)
				visited[neighbor] = true
			}
		}
	}

	return false
}

func (al *AdjacencyList[V]) ConnectedComponentsCount() int {

	visited := make(map[V]bool)
	count := 0

	for node := range *al {

		if visited[node] {
			continue
		}

		count++

		qu := queue.New[V]()
		qu.Enqueue(node)

		for !qu.IsEmpty() {
			current := qu.Dequeue()

			visited[current] = true

			neighbors := (*al)[current]

			for _, neighbor := range neighbors {
				if !visited[neighbor] {
					qu.Enqueue(neighbor)
					visited[neighbor] = true
				}
			}

		}
	}

	return count
}

func (al *AdjacencyList[V]) LargestComponentSize() int {
	visited := make(map[V]bool)
	largest := 0

	for node := range *al {
		if visited[node] {
			continue
		}

		size := 0

		qu := queue.New[V]()
		qu.Enqueue(node)

		for !qu.IsEmpty() {
			current := qu.Dequeue()

			size++

			visited[current] = true

			neighbors := (*al)[current]

			for _, neighbor := range neighbors {
				if !visited[neighbor] {
					qu.Enqueue(neighbor)
					visited[neighbor] = true
				}
			}
		}

		if size > largest {
			largest = size
		}
	}

	return largest
}
