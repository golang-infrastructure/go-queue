package queue

//// ArrayQueue 基于数组实现的队列
//type ArrayQueue[T any] struct {
//	// 队列头的位置
//	headIndex int
//	// 队列尾的位置
//	tailIndex int
//	// 底层的切片
//	slice []T
//}
//
//var _ Queue[any] = &ArrayQueue[any]{}
//
//// NewArrayQueue 创建一个基于数组的队列，其底层其实就是一个长度限制很大的循环队列
//func NewArrayQueue[T any]() *ArrayQueue[T] {
//	return &ArrayQueue[T]{
//		slice: make([]T, 0),
//	}
//}
//
//// TODO 2022-10-29 23:10:47 检查数组大小自动扩容
//func (x *ArrayQueue[T]) computeNextIndex() int {
//	nextIndex :=
//}
//
//func (x *ArrayQueue[T]) Put(values ...T) error {
//	for _, value := range values {
//
//		x.slice = append(x.slice, value)
//		x.headIndex++
//	}
//	return nil
//}
//
//func (x *ArrayQueue[T]) Take() T {
//	e, err := x.TakeE()
//	if err != nil {
//		return nil
//	}
//	return e
//}
//
//func (x *ArrayQueue[T]) TakeE() (T, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) Peek() T {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) PeekE() (T, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) IsEmpty() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) IsNotEmpty() bool {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) Len() int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) Clear() error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) String() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) MarshalJSON() ([]byte, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (x *ArrayQueue[T]) UnmarshalJSON(bytes []byte) error {
//	//TODO implement me
//	panic("implement me")
//}
