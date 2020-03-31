package models

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	SubscriberDefault
	Username string
	WS       *websocket.Conn
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

func (u *User) OnReceive(msg string) {
	message := Message{
		Type: MTMessage,
		Data: fmt.Sprintf("%s: %s", u.GetID(), msg),
	}
	if err := u.WS.WriteJSON(message); err != nil {
		log.Printf("ws send message err: %v", err)
	}
}

func (u *User) GetID() string {
	return u.Username
}

func (u *User) SetWS(ws *websocket.Conn) {
	u.WS = ws
}
