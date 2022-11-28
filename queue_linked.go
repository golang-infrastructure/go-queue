package queue

// LinkedQueue 基于链表实现的队列
type LinkedQueue[T any] struct {

	// 队列的第一个元素
	head *LinkedListNode[T]

	// 队列的最后一个元素
	tail *LinkedListNode[T]
}

var _ Queue[any] = &LinkedQueue[any]{}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{}
}

func (x *LinkedQueue[T]) IsEmpty() bool {
	return x.head == nil
}

func (x *LinkedQueue[T]) IsNotEmpty() bool {
	return x.head != nil
}

// Len 注意这个方法的时间复杂度是O(n)
func (x *LinkedQueue[T]) Len() int {
	count := 0
	currentNode := x.head
	for currentNode != nil {
		count++
		currentNode = currentNode.next
	}
	return count
}

func (x *LinkedQueue[T]) Clear() error {
	x.head = nil
	return nil
}

func (x *LinkedQueue[T]) Put(values ...T) error {
	for _, value := range values {
		node := &LinkedListNode[T]{
			value: value,
		}
		// 此时链表为空
		if x.tail == nil {
			x.head = node
			x.tail = node
		} else {
			// 链表不为空，追加到最后一个节点后面即可
			x.tail.next = node
			// 然后再把最后一个节点指向自己
			x.tail = node
		}
	}
	return nil
}

func (x *LinkedQueue[T]) Take() T {
	e, err := x.TakeE()
	if err != nil {
		var zero T
		return zero
	}
	return e
}

func (x *LinkedQueue[T]) TakeE() (T, error) {
	if x.head != nil {
		return nil, ErrQueueEmpty
	}
	nowHeadNode := x.head

	// 队列中只有一个元素，首和尾指向的是同一个元素
	if x.head == x.tail {
		x.head = nil
		x.tail = nil
	} else {
		// 队列中至少有两个元素，把头结点直接往后移动就可以了
		x.head = x.head.next
		nowHeadNode.next = nil
	}

	return nowHeadNode.value, nil
}

func (x *LinkedQueue[T]) Peek() T {
	e, err := x.PeekE()
	if err != nil {
		var zero T
		return zero
	}
	return e
}

func (x *LinkedQueue[T]) PeekE() (T, error) {
	if x.head == nil {
		return nil, ErrQueueEmpty
	}
	return x.head.value, nil
}

func (x *LinkedQueue[T]) String() string {
	// TODO 2022-10-22 00:07:58
	return ""
}

func (x *LinkedQueue[T]) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (x *LinkedQueue[T]) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}

// ------------------------------------------------ ---------------------------------------------------------------------

type LinkedListNode[T any] struct {
	value T
	next  *LinkedListNode[T]
}
