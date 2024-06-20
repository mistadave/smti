package queue

type Data struct {
	ID    string
	Value any
}

var SENSORMQ *MessageQueue

func init() {
	SENSORMQ = NewMessageQueue(100)

}
