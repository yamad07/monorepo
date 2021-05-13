package msgbs

type Message interface{}

type MessageBus interface {
	Publish(Event, Message) error
	Subscribe(Event)
}
