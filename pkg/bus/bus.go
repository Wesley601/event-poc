package bus

import (
	"github.com/nats-io/nats.go"
)

type EventBus interface {
	Publish(topic string, payload []byte) error
	Subscribe(topic string) (ch chan *nats.Msg, err error)
	Close()
}

type EventBusImpl struct {
	nc  *nats.Conn
	sub *nats.Subscription
}

func New() (*EventBusImpl, error) {
	nc, err := nats.Connect("nats://poc_nats:4222")
	if err != nil {
		return nil, err
	}

	return &EventBusImpl{
		nc: nc,
	}, err
}

func (e *EventBusImpl) Publish(topic string, payload []byte) error {
	return e.nc.Publish(topic, payload)
}

func (e *EventBusImpl) Subscribe(topic string) (ch chan *nats.Msg, err error) {
	// Channel Subscriber
	ch = make(chan *nats.Msg)
	e.sub, err = e.nc.ChanSubscribe("foo", ch)

	if err != nil {
		return nil, err
	}
	return ch, err
}

func (e *EventBusImpl) Close() {
	e.nc.Close()
}
