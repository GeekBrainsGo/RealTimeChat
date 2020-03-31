package factory

// Chater handles chat interactions.
type Chater interface {
	Send(msg string) error      // Used to send messages.
	OnReceive(msg string) error // Handles recieve from ... user or bot.
	Leave() error               // leaving chat?
}

// ClientDefault implements chater interface.
type ClientDefault struct {
	Token string
}

// Send sends client message.
func (c *ClientDefault) Send(string) error {
	panic("Not implemented")
}

// OnReceive do something on receive.
func (c *ClientDefault) OnReceive(string) error {
	panic("Not implemented")
}

// Leave stands for leaving chat.
func (c *ClientDefault) Leave() error {
	panic("Not implemented")
}
