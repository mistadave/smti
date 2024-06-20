package queue

type MessageQueue struct {
	channel chan Data
}

func NewMessageQueue(size int) *MessageQueue {
	return &MessageQueue{
		channel: make(chan Data, size),
	}
}

func (mq *MessageQueue) Put(message Data) {
	mq.channel <- message
}

func (mq *MessageQueue) Get() Data {
	return <-mq.channel
}
