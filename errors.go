package queue

import "errors"

var ErrQueueEmpty = errors.New("queue is empty")

var ErrNotSupportOperation = errors.New("not support")
