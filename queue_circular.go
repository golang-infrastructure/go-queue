package queue

//// 循环队列类似于ArrayQueue，不同发的是循环队列的容量是固定的，而ArrayQueue会自动扩容
//
//type CircularQueue[V any] struct {
//	slice     []V
//	headIndex int
//	tailIndex int
//}
//
//var _ Queue[any] = &CircularQueue[any]{}
//
//func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
//	return &CircularQueue[T]{
//		slice:     make([]T, capacity),
//		headIndex: -1,
//		tailIndex: -1,
//	}
//}
//
//func (x *CircularQueue[V]) Put(values ...V) error {
//	// 循环队列一次只能放一个元素，放多个不支持
//	if len(values) != 1 {
//		return ErrNotSupportOperation
//	}
//	// 计算下一个元素应该放入的位置
//	nextTailIndex := (x.tailIndex + 1) % len(x.slice)
//	if nextTailIndex == x.headIndex {
//		return ErrQueueFull
//	}
//	x.slice[nextTailIndex] = values[0]
//	x.tailIndex = nextTailIndex
//	return nil
//}
//
//func (x *CircularQueue[V]) Take() V {
//
//}
//
//func (x *CircularQueue[V]) TakeE() (V, error) {
//	nextHeadIndex := (x.headIndex + len(x.slice) + 1) % len(x.slice)
//	if nextHeadIndex == x.tailIndex {
//		var zero V
//		return zero, ErrQueueEmpty
//	}
//}
//
//func (x *CircularQueue[V]) Peek() V {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *CircularQueue[V]) PeekE() (V, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *CircularQueue[V]) IsEmpty() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *CircularQueue[V]) IsNotEmpty() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *CircularQueue[V]) Len() int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *CircularQueue[V]) Clear() error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *CircularQueue[V]) String() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *CircularQueue[V]) MarshalJSON() ([]byte, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *CircularQueue[V]) UnmarshalJSON(bytes []byte) error {
//	//TODO implement me
//	panic("implement me")
//}
