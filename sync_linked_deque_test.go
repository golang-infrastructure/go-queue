package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyncLinkedDeque_Clear(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedDeque_IsEmpty(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedDeque_IsNotEmpty(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedDeque_Len(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedDeque_Peek(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, "111", queue.Peek())
}

func TestSyncLinkedDeque_PeekE(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	s, err := queue.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, "111", s)
}

func TestSyncLinkedDeque_Put(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedDeque_Take(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	value := queue.Take()
	assert.Equal(t, "111", value)
}

func TestSyncLinkedDeque_TakeE(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	value, err := queue.TakeE()
	assert.Nil(t, err)
	assert.Equal(t, "111", value)
}

func TestNewSyncLinkedDeque(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	assert.NotNil(t, queue)
}

func TestSyncLinkedDeque_String(t *testing.T) {
	queue := NewSyncLinkedDeque[string]()
	err := queue.Put("1")
	assert.Nil(t, err)
	s := queue.String()
	assert.Equal(t, "[]string{\"1\"}", s)
}

func TestSyncLinkedDeque_MarshalJSON(t *testing.T) {
	queue := NewSyncLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	jsonBytes, err := queue.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "[1]", string(jsonBytes))
}

func TestSyncLinkedDeque_UnmarshalJSON(t *testing.T) {
	queue := NewSyncLinkedDeque[int]()
	jsonString := "[1]"
	err := queue.UnmarshalJSON([]byte(jsonString))
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Take())
	assert.True(t, queue.IsEmpty())
}

func TestSyncLinkedDeque_PeekLast(t *testing.T) {
	queue := NewSyncLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.PeekLast())
}

func TestSyncLinkedDeque_PeekLastE(t *testing.T) {
	queue := NewSyncLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	v, err := queue.PeekLastE()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
}

func TestSyncLinkedDeque_PutFirst(t *testing.T) {
	queue := NewSyncLinkedDeque[int]()
	err := queue.PutFirst(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Peek())
}

func TestSyncLinkedDeque_TakeLast(t *testing.T) {
	queue := NewSyncLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.TakeLast())
}

func TestSyncLinkedDeque_TakeLastE(t *testing.T) {
	queue := NewSyncLinkedDeque[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	v, err := queue.TakeLastE()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
}
