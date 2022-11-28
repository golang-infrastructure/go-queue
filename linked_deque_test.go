package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedDeque_Clear(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedDeque_IsEmpty(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedDeque_IsNotEmpty(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedDeque_Len(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedDeque_Peek(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, "111", queue.Peek())
}

func TestLinkedDeque_PeekE(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	s, err := queue.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, "111", s)
}

func TestLinkedDeque_Put(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedDeque_Take(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	value := queue.Take()
	assert.Equal(t, "111", value)
}

func TestLinkedDeque_TakeE(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	value, err := queue.TakeE()
	assert.Nil(t, err)
	assert.Equal(t, "111", value)
}

func TestNewLinkedDeque(t *testing.T) {
	queue := NewLinkedDeque[string]()
	assert.NotNil(t, queue)
}

func TestLinkedDeque_String(t *testing.T) {
	queue := NewLinkedDeque[string]()
	err := queue.Put("1")
	assert.Nil(t, err)
	s := queue.String()
	assert.Equal(t, "[]string{\"1\"}", s)
}

func TestLinkedDeque_MarshalJSON(t *testing.T) {
	queue := NewLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	jsonBytes, err := queue.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "[1]", string(jsonBytes))
}

func TestLinkedDeque_UnmarshalJSON(t *testing.T) {
	queue := NewLinkedDeque[int]()
	jsonString := "[1]"
	err := queue.UnmarshalJSON([]byte(jsonString))
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Take())
	assert.True(t, queue.IsEmpty())
}

func TestLinkedDeque_PeekLast(t *testing.T) {
	queue := NewLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.PeekLast())
}

func TestLinkedDeque_PeekLastE(t *testing.T) {
	queue := NewLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	v, err := queue.PeekLastE()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
}

func TestLinkedDeque_PutFirst(t *testing.T) {
	queue := NewLinkedDeque[int]()
	err := queue.PutFirst(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Peek())
}

func TestLinkedDeque_TakeLast(t *testing.T) {
	queue := NewLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.TakeLast())
}

func TestLinkedDeque_TakeLastE(t *testing.T) {
	queue := NewLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	v, err := queue.TakeLastE()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
}
