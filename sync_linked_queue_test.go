package queue


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyncLinkedQueue_Clear(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedQueue_IsEmpty(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedQueue_IsNotEmpty(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedQueue_Len(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedQueue_Peek(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, "111", queue.Peek())
}

func TestSyncLinkedQueue_PeekE(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	s, err := queue.PeekE()
	assert.Nil(t, err)
	assert.Equal(t, "111", s)
}

func TestSyncLinkedQueue_Put(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Len())
	err = queue.Clear()
	assert.Nil(t, err)
	assert.Equal(t, 0, queue.Len())
	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsNotEmpty())
}

func TestSyncLinkedQueue_Take(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	value := queue.Take()
	assert.Equal(t, "111", value)
}

func TestSyncLinkedQueue_TakeE(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("111")
	assert.Nil(t, err)
	value, err := queue.TakeE()
	assert.Nil(t, err)
	assert.Equal(t, "111", value)
}

func TestNewSyncLinkedQueue(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	assert.NotNil(t, queue)
}

func TestSyncLinkedQueue_String(t *testing.T) {
	queue := NewSyncLinkedQueue[string]()
	err := queue.Put("1")
	assert.Nil(t, err)
	s := queue.String()
	assert.Equal(t, "[]string{\"1\"}", s)
}

func TestSyncLinkedQueue_MarshalJSON(t *testing.T) {
	queue := NewSyncLinkedQueue[int]()
	err := queue.Put(1)
	assert.Nil(t, err)
	jsonBytes, err := queue.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "[1]", string(jsonBytes))
}

func TestSyncLinkedQueue_UnmarshalJSON(t *testing.T) {
	queue := NewSyncLinkedQueue[int]()
	jsonString := "[1]"
	err := queue.UnmarshalJSON([]byte(jsonString))
	assert.Nil(t, err)
	assert.Equal(t, 1, queue.Take())
	assert.True(t, queue.IsEmpty())
}

