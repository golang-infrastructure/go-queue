package queue

import "encoding/json"

// Queue 队列的接口定义
type Queue[V any] interface {

	// Put 往队列中放元素，如果放失败会返回error
	Put(values ...V) error

	// Take 从队列中取出一个元素，如果队列为空，则返回对应类型的空值
	Take() V

	// TakeE 从队列中取元素，如果队列为空会返回error
	TakeE() (V, error)

	// Peek 取出队列头部的元素，如果队列为空会返回对应类型的零值
	Peek() V

	// PeekE 取出队列头部的元素，如果队列为空会返回error
	PeekE() (V, error)

	// IsEmpty 判断队列是否空
	IsEmpty() bool

	// IsNotEmpty 队列是否非空
	IsNotEmpty() bool

	// Len 队列中元素的个数
	Len() int

	// Clear 清空队列中的元素
	Clear() error

	// String 一般用于日志之类的场景，需要转化为字符串打印
	String() string

	// Marshaler 队列都应该是可以被序列化的
	json.Marshaler
	json.Unmarshaler
}
