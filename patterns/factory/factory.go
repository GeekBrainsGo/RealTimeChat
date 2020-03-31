package factory

import "errors"

// NewChater returns new clienter interface.
func NewChater(token, typ string) (Chater, error) {
	switch typ {
	case "user":
		return NewUser(token), nil
	case "bot":
		return NewBot(token), nil
	default:
		return nil, errors.New("unknown client type")
	}
}
