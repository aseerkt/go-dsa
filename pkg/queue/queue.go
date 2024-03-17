package queue

type node[V comparable] struct {
	value V
	prev  *node[V]
}

type Queue[V comparable] struct {
	first  *node[V]
	last   *node[V]
	length int
}

func New[V comparable]() *Queue[V] {
	return &Queue[V]{nil, nil, 0}
}

func (q *Queue[V]) Enqueue(v V) {
	n := &node[V]{v, nil}

	if q.last == nil {
		q.last, q.first = n, n
	} else {
		q.last.prev, q.last = n, n
	}
	q.length++
}

func (q *Queue[V]) Dequeue() V {
	if q.first == nil {
		var value V
		return value
	}

	n := q.first

	if q.length == 1 {
		q.first, q.last = nil, nil
	} else {
		q.first = n.prev
	}

	q.length--
	return n.value
}

func (q *Queue[V]) IsEmpty() bool {
	return q.length == 0
}
