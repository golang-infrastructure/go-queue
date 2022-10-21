package queue

import "context"

// Queue 队列的接口定义
type Queue[T any] interface {
	Put(value ...T) (err error)
	BPut(ctx context.Context, value T) (err error)

	Take() (value T, err error)
	BTake(ctx context.Context) (value T, err error)

	Peek() (value T, err error)

	IsEmpty() bool
	IsNotEmpty() bool
	Size() int
	Clear()

	String() string
}
