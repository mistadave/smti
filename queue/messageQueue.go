package queue

type MessageQueue struct {
	Channel chan Data
}

func NewMessageQueue(size int) *MessageQueue {
	return &MessageQueue{
		Channel: make(chan Data),
	}
}
