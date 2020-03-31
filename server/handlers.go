package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/phadeyev/RealTimeChat/models"
)

func (serv *Server) WShandler(w http.ResponseWriter, r *http.Request) {

	ws, err := serv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			<-time.After(5 * time.Second)
			msg := Message{
				Type: MTPing,
			}
			if err := ws.WriteJSON(msg); err != nil {
				log.Printf("ws send ping err: %v", err)
				break
			}
		}
	}()
	client := models.NewUser(uuid.New().String())
	fmt.Println("client connected: ", client.GetID())
	client.SetWS(ws)
	ch, err := serv.publisher.GetChannel("one")
	if err != nil {
		fmt.Println(err)
	}
	ch.Subscribe(client)
	for {
		msg := Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, 1001) {
				log.Println("ws msg read err: %v", err)
			}
			break
		}

		if msg.Type == MTPong {
			continue
		}

		if msg.Type == MTMessage {
			serv.submutex.Lock()
			serv.publisher.Send(msg.Data, "one")
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
