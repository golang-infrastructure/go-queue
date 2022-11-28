package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedQueue_Clear(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedQueue_IsEmpty(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedQueue_IsNotEmpty(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedQueue_Len(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedQueue_Peek(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, "111", queue.Peek())
}

func TestLinkedQueue_PeekE(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	s, err := queue.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, "111", s)
}

func TestLinkedQueue_Put(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestLinkedQueue_Take(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	value := queue.Take()
	assert.Equal(t, "111", value)
}

func TestLinkedQueue_TakeE(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	value, err := queue.TakeE()
	assert.Nil(t, err)
	assert.Equal(t, "111", value)
}

func TestNewLinkedQueue(t *testing.T) {
	queue := NewLinkedQueue[string]()
	assert.NotNil(t, queue)
}

func TestLinkedQueue_String(t *testing.T) {
	queue := NewLinkedQueue[string]()
	err := queue.Put("1")
	assert.Nil(t, err)
	s := queue.String()
	assert.Equal(t, "[]string{\"1\"}", s)
}

func TestLinkedQueue_MarshalJSON(t *testing.T) {
	queue := NewLinkedQueue[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	jsonBytes, err := queue.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "[1]", string(jsonBytes))
}

func TestLinkedQueue_UnmarshalJSON(t *testing.T) {
	queue := NewLinkedQueue[int]()
	jsonString := "[1]"
	err := queue.UnmarshalJSON([]byte(jsonString))
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Take())
	assert.True(t, queue.IsEmpty())
}
