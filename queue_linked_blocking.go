package queue

import "context"

// 链表阻塞队列

// TODO 2022-10-29 22:12:06 阻塞队列基于channel实现，等实现了能够自动扩容的channel再来实现这个

type BlockingLinkedQueue[T any] struct {
}

var _ BlockingQueue[any] = &BlockingLinkedQueue[any]{}

func NewBlockingLinkedQueue[T any]() *BlockingLinkedQueue[T] {
	return &BlockingLinkedQueue[T]{}
}

func (x *BlockingLinkedQueue[T]) Put(values ...T) error {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) Take() T {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) TakeE() (T, error) {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) Peek() T {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) PeekE() (T, error) {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) IsNotEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) Clear() error {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) String() string {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) BPut(ctx context.Context, value T) error {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) BTake(ctx context.Context) T {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) BTakeE(ctx context.Context) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) BPeek(ctx context.Context) T {
	//TODO implement me
	panic("implement me")
}

func (x *BlockingLinkedQueue[T]) BPeekE(ctx context.Context) (T, error) {
	//TODO implement me
	panic("implement me")
}
