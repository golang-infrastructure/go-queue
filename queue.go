package queue

// Queue 队列的接口定义
type Queue[T any] interface {
	Put(values ...T) (err error)
	Take() (value T, err error)
	Peek() (value T, err error)
	IsEmpty() bool
	IsNotEmpty() bool
	Len() int
	Clear() error
	String() string
}
