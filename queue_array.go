package queue

// ArrayQueue 基于数组实现的队列
type ArrayQueue[T any] struct {
	headIndex int
	tailIndex int
	slice     []T
}

//var _ Queue[any] = &ArrayQueue[any]{}

// NewArrayQueue 创建一个基于数组的队列，其底层其实就是一个长度限制很大的循环队列
func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{
		slice: make([]T, 0),
	}
}
