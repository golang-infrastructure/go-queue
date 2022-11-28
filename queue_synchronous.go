package queue

//import (
//	"context"
//	"time"
//)
//
//// SynchronousQueue 同步队列，类似于无缓冲的channel，不同的是在channel的基础上增加了超时的支持
//// 当有多个Take被阻塞时，并发的公平性与Go内置的Channel保持一致，为先来的先取到
//type SynchronousQueue[T any] struct {
//	noBuff chan T
//}
//
//var _ BlockingQueue[any] = &SynchronousQueue[any]{}
//
//func NewSynchronousQueue[T any]() *SynchronousQueue[T] {
//	return &SynchronousQueue[T]{
//		noBuff: make(chan T, 0),
//	}
//}
//
//func (x *SynchronousQueue[T]) Put(values ...T) error {
//	return ErrNotSupportOperation
//}
//
//func (x *SynchronousQueue[T]) Take() T {
//	return nil
//}
//
//func (x *SynchronousQueue[T]) TakeE() (T, error) {
//	return nil, ErrNotSupportOperation
//}
//
//func (x *SynchronousQueue[T]) Peek() T {
//	return nil
//}
//
//func (x *SynchronousQueue[T]) PeekE() (T, error) {
//	return nil, ErrNotSupportOperation
//}
//
//func (x *SynchronousQueue[T]) IsEmpty() bool {
//	return len(x.noBuff) == 0
//}
//
//func (x *SynchronousQueue[T]) IsNotEmpty() bool {
//	return len(x.noBuff) != 0
//}
//
//func (x *SynchronousQueue[T]) Len() int {
//	return len(x.noBuff)
//}
//
//func (x *SynchronousQueue[T]) Clear() error {
//	// 尝试以极短的超时时间读取一个元素
//	select {
//	case <-x.noBuff:
//	case <-time.After(time.Millisecond * 10):
//	}
//	return nil
//}
//
//// BPut 往同步队列中放入一个元素，会卡住直到被取出或者超时
//func (x *SynchronousQueue[T]) BPut(ctx context.Context, value T) error {
//	select {
//	case x.noBuff <- value:
//		return nil
//	case <-ctx.Done():
//		return context.DeadlineExceeded
//	}
//}
//
//func (x *SynchronousQueue[T]) BTake(ctx context.Context) T {
//	e, err := x.BTakeE(ctx)
//	if err != nil {
//		return nil
//	}
//	return e
//}
//
//// BTakeE 尝试从同步队列中取出一个元素，会阻塞住直到获取到元素或者超时
//func (x *SynchronousQueue[T]) BTakeE(ctx context.Context) (T, error) {
//	select {
//	case v := <-x.noBuff:
//		return v, nil
//	case <-ctx.Done():
//		return nil, context.DeadlineExceeded
//	}
//}
//
//func (x *SynchronousQueue[T]) BPeek(ctx context.Context) T {
//	return nil
//}
//
//func (x *SynchronousQueue[T]) BPeekE(ctx context.Context) (T, error) {
//	return nil, ErrNotSupportOperation
//}
//
//func (x *SynchronousQueue[T]) String() string {
//	//TODO implement me
//	panic("implement me")
//}
