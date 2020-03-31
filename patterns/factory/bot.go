package factory

// Bot stands for bot.
type Bot struct {
	ClientDefault
	Token string
}

// NewBot returns new bot.
func NewBot(token string) *Bot {
	return &Bot{Token: token}
}

// Send sends user message.
func (b *Bot) Send(msg string) error {
	return nil
}

// OnReceive do something on receive.
func (b *Bot) OnReceive(msg string) error {
	return nil
}
