package eventchannel

import (
	"fmt"
)

// Channels stands for channel dictionary.
type Channels map[string]*Channel

// Publisher implements basic struct of event channel pattern.
type Publisher struct {
	channels Channels
}

// NewPublisher returns new publisher.
func NewPublisher() *Publisher {
	return &Publisher{
		channels: Channels{},
	}
}

// AddChannel adds new channel.
func (p *Publisher) AddChannel(name string, channel *Channel) {
	p.channels[name] = channel
}

// Broadcast sends message to all channels.
func (p *Publisher) Broadcast(msg string) {
	for _, ch := range p.channels {
		ch.Send(msg)
	}
}

// DeleteChannel removes channel.
func (p *Publisher) DeleteChannel(name string) error {
	if ch, ok := p.channels[name]; ok {
		delete(p.channels, name)
		return ch.UnSubscribeAll()
	}
	return fmt.Errorf("can't find channel with %q name", name)
}

// Send sends message to channels.
func (p *Publisher) Send(msg string, channels ...string) error {
	if len(channels) == 0 {
		p.Broadcast(msg)
		return nil
	}
	for _, ch := range channels {
		c, ok := p.channels[ch]
		if !ok {
			return fmt.Errorf("channel %q can't be found", ch)
		}
		c.Send(msg)
	}
	return nil
}
