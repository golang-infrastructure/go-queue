package queue

import (
	"sync"
)

// SyncLinkedQueue 线程安全的链表队列
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

func (x *SyncLinkedQueue[T]) Put(values ...T) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Put(values...)
}

func (x *SyncLinkedQueue[T]) Take() T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Take()
}

func (x *SyncLinkedQueue[T]) TakeE() (T, error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.TakeE()
}

func (x *SyncLinkedQueue[T]) Peek() T {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.Peek()
}

func (x *SyncLinkedQueue[T]) PeekE() (T, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.PeekE()
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
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.String()
}

func (x *SyncLinkedQueue[T]) MarshalJSON() ([]byte, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.MarshalJSON()
}

func (x *SyncLinkedQueue[T]) UnmarshalJSON(bytes []byte) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.UnmarshalJSON(bytes)
}
