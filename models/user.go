package models

import (
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	SubscriberDefault
	Username string
	WS       *websocket.Conn
}

func NewUser(username string, ws *websocket.Conn) *User {
	return &User{
		Username: username,
		WS:       ws,
	}
}

func (u *User) OnReceive(msg string) {
	message := Message{
		Type: MTMessage,
		Data: msg,
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
