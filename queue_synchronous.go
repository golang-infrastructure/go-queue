package queue

import "context"

// 同步队列，类似于无缓冲的channel，不同的是在channel的基础上增加了超时的支持

type SynchronousQueue[T any] struct {
	noBuff chan T
}

var _ Queue[any] = &SynchronousQueue[any]{}

func NewSynchronousQueue[T any]() *SynchronousQueue[T] {
	return &SynchronousQueue[T]{
		noBuff: make(chan T, 0),
	}
}

// BPut 往同步队列中放入一个元素，会卡住直到被取出或者超时
func (x *SynchronousQueue[T]) BPut(ctx context.Context, value T) error {
	select {
	case x.noBuff <- value:
		return nil
	case <-ctx.Done():
		return ErrTimeout
	}
}

// BTake 尝试从同步队列中取出一个元素，会阻塞住直到获取到元素或者超时
func (x *SynchronousQueue[T]) BTake(ctx context.Context) (T, error) {
	select {
	case v := <-x.noBuff:
		return v, nil
	case <-ctx.Done():
		return nil, ErrTimeout
	}
}

func (x *SynchronousQueue[T]) Put(value ...T) (err error) {
	return ErrNotSupportOperation
}

func (x *SynchronousQueue[T]) Take() (value T, err error) {
	return nil, ErrNotSupportOperation
}

func (x *SynchronousQueue[T]) Peek() (value T, err error) {
	return nil, ErrNotSupportOperation
}

func (x *SynchronousQueue[T]) IsEmpty() bool {
	return true
}

func (x *SynchronousQueue[T]) IsNotEmpty() bool {
	return false
}

func (x *SynchronousQueue[T]) Size() int {
	return 0
}

func (x *SynchronousQueue[T]) Clear() {
}

func (x *SynchronousQueue[T]) String() string {
	return ""
}
