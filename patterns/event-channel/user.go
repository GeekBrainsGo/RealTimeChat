package eventchannel

import (
	"chat/message"
	"log"

	"github.com/gorilla/websocket"
)

// User stands for humble user.
type User struct {
	SubscriberDefault
	Name string
	conn *websocket.Conn
}

// NewUser returns new user.
func NewUser(name string, ws *websocket.Conn) *User {
	return &User{Name: name, conn: ws}
}

// OnReceive do something on receive.
func (u *User) OnReceive(msg string) {
	m := message.Message{
		Type: message.MTMessage,
		Data: msg,
	}
	if err := u.conn.WriteJSON(&m); err != nil {
		log.Println(err)
	}
}

// GetID returns user id.
func (u *User) GetID() string {
	return u.Name
}
