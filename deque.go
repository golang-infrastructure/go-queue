package queue

// Deque 双端队列
type Deque[T any] interface {

	// Queue 双端队列继承了Queue的API，同时会在此基础上新增一些API
	Queue[T]

	// TODO 2022-10-29 19:35:50

}
