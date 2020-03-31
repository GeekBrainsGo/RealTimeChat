package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/phadeyev/RealTimeChat/models"
)

func (serv *Server) WShandler(w http.ResponseWriter, r *http.Request) {

	ws, err := serv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	go SendPing(ws)

	client := models.NewUser(uuid.New().String(), ws)
	fmt.Println("client connected: ", client.GetID())

	ch, err := serv.publisher.GetChannel("one")
	if err != nil {
		fmt.Println(err)
	}
	ch.Subscribe(client)
	for {
		msg := models.Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, 1001) {
				log.Println("ws msg read err: %v", err)
			}
			break
		}

		if msg.Type == models.MTPong {
			continue
		}

		if msg.Type == models.MTMessage {
			serv.submutex.Lock()
			serv.publisher.Send(fmt.Sprintf("%s: %s", client.GetID(), msg.Data), "one")
			serv.submutex.Unlock()
		}
	}
	defer func() {
		ws.Close()
		serv.submutex.Lock()
		ch, err := serv.publisher.GetChannel("one")
		if err != nil {
			fmt.Println(err)
		}
		ch.UnSubscribe(client)
		serv.submutex.Unlock()
		fmt.Println("client disconnected: ", client.GetID())
	}()

}
