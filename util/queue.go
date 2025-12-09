package util

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return make(Queue[T], 0)
}

func (q *Queue[T]) Push(items ...T) {
	*q = append(*q, items...)
}

func (q *Queue[T]) Pop() T {
	item := q.Peek()
	*q = (*q)[:len(*q)-1]
	return item
}

func (q *Queue[T]) Peek() T {
	return (*q)[len(*q)-1]
}

func (q *Queue[T]) Seq(yield func(T) bool) {
	for len(*q) > 0 {
		if !yield(q.Pop()) {
			return
		}
	}
}
