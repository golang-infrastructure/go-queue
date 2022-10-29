package queue

import "context"

// BlockingQueue 阻塞队列所拥有的的方法的定义
type BlockingQueue[T any] interface {
	Queue[T]

	BPut(ctx context.Context, value T) error

	BTake(ctx context.Context) T
	BTakeE(ctx context.Context) (T, error)

	// BPeek 取出队列头部的元素，如果队列为空会返回对应类型的零值
	BPeek(ctx context.Context) T

	// BPeekE 取出队列头部的元素，如果队列为空会返回error
	BPeekE(ctx context.Context) (T, error)
}
