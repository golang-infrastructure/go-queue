package queue

import (
	"context"
	"sync"
)

// 阻塞队列

type SyncLinkedQueue[T any] struct {
	queue *LinkedQueue[T]
	lock  sync.RWMutex
}

var _ Queue[any] = &SyncLinkedQueue[any]{}

func NewSyncLinkedQueue[T any]() *SyncLinkedQueue[T] {
	return &SyncLinkedQueue[T]{
		queue: NewLinkedQueue[T](),
		lock:  sync.RWMutex{},
	}
}

func (x *SyncLinkedQueue[T]) Put(values ...T) (err error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Put(values...)
}

func (x *SyncLinkedQueue[T]) BPut(ctx context.Context, value T) (err error) {
	return ErrNotSupportOperation
}

func (x *SyncLinkedQueue[T]) Take() (value T, err error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Take()
}

func (x *SyncLinkedQueue[T]) BTake(ctx context.Context) (value T, err error) {
	return nil, ErrNotSupportOperation
}

func (x *SyncLinkedQueue[T]) Peek() (value T, err error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Peek()
}

func (x *SyncLinkedQueue[T]) IsEmpty() bool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.IsEmpty()
}

func (x *SyncLinkedQueue[T]) IsNotEmpty() bool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.IsNotEmpty()
}

func (x *SyncLinkedQueue[T]) Len() int {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.Len()
}

func (x *SyncLinkedQueue[T]) Clear() error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Clear()
}

func (x *SyncLinkedQueue[T]) String() string {
	//TODO implement me
	panic("implement me")
}
