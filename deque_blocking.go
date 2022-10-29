package queue

// BlockingDeque 阻塞双向队列
type BlockingDeque[T any] interface {
	Deque[T]
}
