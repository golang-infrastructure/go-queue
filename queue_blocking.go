package queue

import "context"

type BlockingQueue[T any] interface {
	Queue[T]
	BPut(ctx context.Context, value T) (err error)
	BTake(ctx context.Context) (value T, err error)
}
