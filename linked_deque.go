package queue

import (
	"encoding/json"
	"fmt"
)

// ------------------------------------------------ LinkedDeque --------------------------------------------------------

// LinkedDeque 基于链表实现的队列
type LinkedDeque[T any] struct {

	// 队列的第一个元素
	head *LinkedDequeNode[T]

	// 队列的最后一个元素
	tail *LinkedDequeNode[T]
}

var _ Deque[any] = &LinkedDeque[any]{}

func NewLinkedDeque[T any]() *LinkedDeque[T] {
	return &LinkedDeque[T]{}
}

func (x *LinkedDeque[T]) IsEmpty() bool {
	return x.head == nil
}

func (x *LinkedDeque[T]) IsNotEmpty() bool {
	return x.head != nil
}

// Len 注意这个方法的时间复杂度是O(n)
func (x *LinkedDeque[T]) Len() int {
	count := 0
	currentNode := x.head
	for currentNode != nil {
		count++
		currentNode = currentNode.next
	}
	return count
}

func (x *LinkedDeque[T]) Clear() error {
	x.head = nil
	return nil
}

func (x *LinkedDeque[T]) Put(values ...T) error {
	for _, value := range values {
		node := &LinkedDequeNode[T]{
			value: value,
		}
		// 此时链表为空
		if x.tail == nil {
			x.head = node
			x.tail = node
		} else {
			// 链表不为空，追加到最后一个节点后面即可
			x.tail.next = node
			node.prev = x.tail
			// 然后再把最后一个节点指向自己
			x.tail = node
		}
	}
	return nil
}

func (x *LinkedDeque[T]) Take() T {
	e, err := x.TakeE()
	if err != nil {
		var zero T
		return zero
	}
	return e
}

func (x *LinkedDeque[T]) TakeE() (T, error) {
	if x.head == nil {
		var zero T
		return zero, ErrQueueEmpty
	}
	nowHeadNode := x.head

	// 队列中只有一个元素，首和尾指向的是同一个元素
	if x.head == x.tail {
		x.head = nil
		x.tail = nil
	} else {
		// 队列中至少有两个元素，把头结点直接往后移动就可以了
		x.head = x.head.next
		// 把关联关系断掉
		x.head.prev = nil
		nowHeadNode.next = nil
	}

	return nowHeadNode.value, nil
}

func (x *LinkedDeque[T]) Peek() T {
	e, err := x.PeekE()
	if err != nil {
		var zero T
		return zero
	}
	return e
}

func (x *LinkedDeque[T]) PeekE() (T, error) {
	if x.head == nil {
		var zero T
		return zero, ErrQueueEmpty
	}
	return x.head.value, nil
}

func (x *LinkedDeque[T]) PutFirst(values ...T) error {
	for _, value := range values {
		node := &LinkedDequeNode[T]{
			value: value,
		}
		// 此时链表为空
		if x.head == nil {
			x.head = node
			x.tail = node
		} else {
			// 链表不为空，放到head前面
			node.next = x.head
			x.head.prev = node
			// 然后把head指针往前移动
			x.head = node
		}
	}
	return nil
}

func (x *LinkedDeque[T]) TakeLast() T {
	value, _ := x.TakeLastE()
	return value
}

func (x *LinkedDeque[T]) TakeLastE() (T, error) {
	if x.tail == nil {
		var zero T
		return zero, ErrQueueEmpty
	}
	nowHeadNode := x.tail

	// 队列中只有一个元素，首和尾指向的是同一个元素
	if x.head == x.tail {
		x.head = nil
		x.tail = nil
	} else {
		// 队列中至少有两个元素，把尾部的节点往前移动
		x.tail.prev = nil
		if x.tail.prev != nil {
			x.tail.prev.next = nil
		}
		// 尾指针往前移动一个位置
		x.tail = x.tail.prev
	}

	return nowHeadNode.value, nil
}

func (x *LinkedDeque[T]) PeekLast() T {
	v, _ := x.PeekLastE()
	return v
}

func (x *LinkedDeque[T]) PeekLastE() (T, error) {
	if x.tail == nil {
		var zero T
		return zero, ErrQueueEmpty
	}
	return x.tail.value, nil
}

func (x *LinkedDeque[T]) toValueSlice() []T {
	result := make([]T, 0)
	currentNode := x.head
	for currentNode != nil {
		result = append(result, currentNode.value)
		currentNode = currentNode.next
	}
	return result
}

func (x *LinkedDeque[T]) String() string {
	return fmt.Sprintf("%#v", x.toValueSlice())
}

func (x *LinkedDeque[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.toValueSlice())
}

func (x *LinkedDeque[T]) UnmarshalJSON(bytes []byte) error {
	slice := make([]T, 0)
	err := json.Unmarshal(bytes, &slice)
	if err != nil {
		return err
	}
	return x.Put(slice...)
}

// ------------------------------------------------ LinkedDequeNode -----------------------------------------------------

// LinkedDequeNode 双向链表
type LinkedDequeNode[T any] struct {
	// 节点的值
	value T
	// 后继节点
	next *LinkedDequeNode[T]
	// 前驱结点
	prev *LinkedDequeNode[T]
}

// ---------------------------------------------------------------------------------------------------------------------
