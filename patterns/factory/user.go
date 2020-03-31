package factory

// User stands for humble default user.
type User struct {
	ClientDefault
	Token string
}

// NewUser returns new user.
func NewUser(token string) *User {
	return &User{Token: token}
}

// Send sends client message.
func (u *User) Send(msg string) error {
	return nil
}

// OnReceive do something on receive.
func (u *User) OnReceive(msg string) error {
	return nil
}

// Leave stands for leaving chat.
func (u *User) Leave() error {
	return nil
}
