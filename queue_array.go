package queue

// ArrayQueue 基于数组实现的队列
type ArrayQueue[V any] struct {
	headIndex int
	tailIndex int
	slice     []V
}

var _ Queue[any] = &ArrayQueue[any]{}

// NewArrayQueue 创建一个基于数组的队列，其底层其实就是一个长度限制很大的循环队列
func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{
		slice: make([]T, 0),
	}
}

func (x *ArrayQueue[V]) Put(values ...V) error {
	for _, value := range values {
		x.slice = append(x.slice, value)
	}
	return nil
}

func (x *ArrayQueue[V]) Take() V {
	
}

func (x *ArrayQueue[V]) TakeE() (V, error) {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) Peek() V {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) PeekE() (V, error) {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) IsNotEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) Clear() error {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) String() string {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (x *ArrayQueue[V]) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}
