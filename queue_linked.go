package queue

// 基于链表实现的队列

type LinkedQueue[T any] struct {
	head *Node[T]
	tail *Node[T]
}

var _ Queue[any] = &LinkedQueue[any]{}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{}
}

func (x *LinkedQueue[T]) Put(values ...T) (err error) {
	for _, value := range values {
		node := &Node[T]{
			value: value,
		}
		// 此时链表为空
		if x.tail == nil {
			x.head = node
			x.tail = node
		} else {
			// 链表不为空，追加到最后即可
			x.tail.next = node
		}
	}
	return nil
}

func (x *LinkedQueue[T]) Take() (value T, err error) {
	if x.head == nil {
		return nil, err
	}
	value = x.head.value

	// 链表中只有一个元素，还被取走了，此时需要将链表置为空
	if x.head == x.tail {
		x.tail = nil
		x.head = nil
	} else {
		x.head = x.head.next
	}
	return value, nil
}

func (x *LinkedQueue[T]) Peek() (value T, err error) {
	if x.head == nil {
		return nil, nil
	}
	return x.head.value, nil
}

func (x *LinkedQueue[T]) IsEmpty() bool {
	return x.head == nil
}

func (x *LinkedQueue[T]) IsNotEmpty() bool {
	return x.head != nil
}

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

func (x *LinkedQueue[T]) String() string {
	// TODO 2022-10-22 00:07:58
	return ""
}

// ------------------------------------------------ ---------------------------------------------------------------------

type Node[T any] struct {
	value T
	next  *Node[T]
}
