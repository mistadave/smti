package queue

type Data struct {
	ID    string
	Tags  map[string]string
	Value interface{}
}

var SENSORMQ *MessageQueue

func init() {
	SENSORMQ = NewMessageQueue(100)

}
