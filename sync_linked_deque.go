package queue

import (
	"sync"
)

// ------------------------------------------------ SyncLinkedDeque --------------------------------------------------------

// SyncLinkedDeque 基于链表实现的队列
type SyncLinkedDeque[T any] struct {
	lock  sync.RWMutex
	queue *LinkedDeque[T]
}

var _ Deque[any] = &SyncLinkedDeque[any]{}

func NewSyncLinkedDeque[T any]() *SyncLinkedDeque[T] {
	return &SyncLinkedDeque[T]{
		queue: NewLinkedDeque[T](),
		lock:  sync.RWMutex{},
	}
}

func (x *SyncLinkedDeque[T]) IsEmpty() bool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.IsEmpty()
}

func (x *SyncLinkedDeque[T]) IsNotEmpty() bool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.IsNotEmpty()
}

// Len 注意这个方法的时间复杂度是O(n)
func (x *SyncLinkedDeque[T]) Len() int {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.Len()
}

func (x *SyncLinkedDeque[T]) Clear() error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Clear()
}

func (x *SyncLinkedDeque[T]) Put(values ...T) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Put(values...)
}

func (x *SyncLinkedDeque[T]) Take() T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.Take()
}

func (x *SyncLinkedDeque[T]) TakeE() (T, error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.TakeE()
}

func (x *SyncLinkedDeque[T]) Peek() T {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.Peek()
}

func (x *SyncLinkedDeque[T]) PeekE() (T, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.PeekE()
}

func (x *SyncLinkedDeque[T]) PutFirst(values ...T) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.PutFirst(values...)
}

func (x *SyncLinkedDeque[T]) TakeLast() T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.TakeLast()
}

func (x *SyncLinkedDeque[T]) TakeLastE() (T, error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.TakeLastE()
}

func (x *SyncLinkedDeque[T]) PeekLast() T {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.PeekLast()
}

func (x *SyncLinkedDeque[T]) PeekLastE() (T, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.PeekLastE()
}

func (x *SyncLinkedDeque[T]) String() string {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.String()
}

func (x *SyncLinkedDeque[T]) MarshalJSON() ([]byte, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.queue.MarshalJSON()
}

func (x *SyncLinkedDeque[T]) UnmarshalJSON(bytes []byte) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.queue.UnmarshalJSON(bytes)
}

// ------------------------------------------------ SyncLinkedDequeNode -----------------------------------------------------

// SyncLinkedDequeNode 双向链表
type SyncLinkedDequeNode[T any] struct {
	// 节点的值
	value T
	// 后继节点
	next *SyncLinkedDequeNode[T]
	// 前驱结点
	prev *SyncLinkedDequeNode[T]
}

// ---------------------------------------------------------------------------------------------------------------------
