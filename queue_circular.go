package queue

type CircularQueue[T any] struct {
	slice     []T
	headIndex int
	tailIndex int
}

//var _ Queue[any] = &CircularQueue[any]{}

func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	return &CircularQueue[T]{
		slice:     make([]T, capacity),
		headIndex: 0,
		tailIndex: capacity - 1,
	}
}

