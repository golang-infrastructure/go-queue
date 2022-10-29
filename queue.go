package queue

// Queue 队列的接口定义
type Queue[T any] interface {

	// Put 往队列中放元素，如果放失败会返回error
	Put(values ...T) error

	// Take 从队列中取出一个元素，如果队列为空，则返回对应类型的空值
	Take() T

	// TakeE 从队列中取元素，如果队列为空会返回error
	TakeE() (T, error)

	// Peek 取出队列头部的元素，如果队列为空会返回对应类型的零值
	Peek() T

	// PeekE 取出队列头部的元素，如果队列为空会返回error
	PeekE() (T, error)

	// IsEmpty 判断队列是否空
	IsEmpty() bool

	// IsNotEmpty 队列是否非空
	IsNotEmpty() bool

	// Len 队列中元素的个数
	Len() int

	// Clear 清空队列中的元素
	Clear() error

	String() string
}
