package queue

//import "context"
//
//// TransferQueue 在BlockingQueue的基础上增加更多的特性，生产者会一直阻塞直到所放入的元素被某个消费者消费完毕，而不是队列放成功就不管了
//type TransferQueue[T any] interface {
//	BlockingQueue[T]
//
//	GetWaitingConsumerCount() int
//
//	HasWaitingConsumer() int
//
//	Transfer(ctx context.Context, element T, callbackFunc func(element T)) error
//}
