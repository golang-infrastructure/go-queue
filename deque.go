package queue

// Deque 双端队列
type Deque[T any] interface {

	// Queue 双端队列继承了Queue的API，同时会在此基础上新增一些API
	Queue[T]

	// PutFirst 可以往队列头部放值了
	PutFirst(values ...T) error

	// TakeLast 可以从队列尾部获取值
	TakeLast() T
	TakeLastE() (T, error)

	// PeekLast 可以查看队列尾部的值
	PeekLast() T
	PeekLastE() (T, error)
}
