package queue

import "github.com/golang-infrastructure/go-heap"

// PriorityQueue 优先级队列
type PriorityQueue[T any] struct {
	heap *heap.Heap[T]
}

var _ Queue[any] = &PriorityQueue[any]{}

type Comparator[T any] func(a T, b T) int

func NewPriorityQueue[T any](comparator Comparator[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: heap.New[T](heap.Comparator[T](comparator)),
	}
}

func (x *PriorityQueue[T]) Put(values ...T) error {
	x.heap.Push(values...)
	return nil
}

func (x *PriorityQueue[T]) Take() T {
	return x.heap.Pop()
}

func (x *PriorityQueue[T]) TakeE() (T, error) {
	return x.heap.PopE()
}

func (x *PriorityQueue[T]) Peek() T {
	return x.heap.Peek()
}

func (x *PriorityQueue[T]) PeekE() (T, error) {
	return x.heap.PeekE()
}

func (x *PriorityQueue[T]) IsEmpty() bool {
	return x.heap.IsEmpty()
}

func (x *PriorityQueue[T]) IsNotEmpty() bool {
	return x.heap.IsNotEmpty()
}

func (x *PriorityQueue[T]) Len() int {
	return x.heap.Size()
}

func (x *PriorityQueue[T]) Clear() error {
	x.heap.Clear()
	return nil
}

func (x *PriorityQueue[T]) String() string {
	//TODO implement me
	panic("implement me")
}

func (x *PriorityQueue[T]) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (x *PriorityQueue[T]) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}
