package models

import "github.com/gorilla/websocket"

type Subscriber interface {
	OnReceive(msg string)
	GetID() string
}

type SubscriberDefault struct{}

func (SubscriberDefault) OnReceive(string) {
	panic("not implemented")
}
func (SubscriberDefault) GetID(string) {
	panic("not implemented")
}

func (SubscriberDefault) SetWS(ws *websocket.Conn) {
	panic("not implemented")
}
