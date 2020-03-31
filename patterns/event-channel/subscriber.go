package eventchannel

// Subscriber handles channel subscriber interactions.
type Subscriber interface {
	OnReceive(msg string)
	GetID() string
}

// SubscriberDefault implements subscriber interface.
type SubscriberDefault struct{}

// OnReceive does some magic with message.
func (SubscriberDefault) OnReceive(string) {
	panic("not implemented")
}

// GetID returns subscriber's id.
func (SubscriberDefault) GetID() string {
	panic("not implemented")
}
