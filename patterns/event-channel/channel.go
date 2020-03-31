package eventchannel

import (
	"fmt"
)

// Subscribers stands for channel subscribers.
type Subscribers map[string]Subscriber

// Channel handles subscriber-channel interactions.
type Channel struct {
	subscribers Subscribers
}

// NewChannel returns new channel.
func NewChannel() *Channel {
	return &Channel{
		subscribers: Subscribers{},
	}
}

// Send sends message to channel's subcribers.
func (c *Channel) Send(msg string) {
	for _, sub := range c.subscribers {
		sub.OnReceive(msg)
	}
}

// Subscribe subscribes subcriber to channel.
func (c *Channel) Subscribe(sub Subscriber) {
	c.subscribers[sub.GetID()] = sub
}

// UnSubscribe unsubcribes subcriber from channel.
func (c *Channel) UnSubscribe(sub Subscriber) error {
	id := sub.GetID()
	if _, ok := c.subscribers[id]; ok {
		delete(c.subscribers, id)
		return nil
	}
	return fmt.Errorf("can't find subcriber with %q id", id)
}

// UnSubscribeAll unsubcribes all subcribers from channel.
func (c *Channel) UnSubscribeAll() error {
	for _, sub := range c.subscribers {
		return c.UnSubscribe(sub)
	}
	return nil
}
