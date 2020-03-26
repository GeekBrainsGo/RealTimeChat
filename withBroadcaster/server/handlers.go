package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/google/uuid"
)

// ApplyHandlers applies all server handlers.
func (s *Server) ApplyHandlers() {
	s.router.Handle("/*", http.FileServer(http.Dir("./web")))
	s.router.Get("/socket", s.socketHandler)
}

func (s *Server) socketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	out := make(chan string) // outgoing client message
	go clientWriter(ws, out)

	cli := &client{out, uuid.New().String()}
	entering <- cli
	defer func() { leaving <- cli }()

	for {
		msg := Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			log.Println(err)
			return
		}
		if msg.Type == MTPong {
			continue
		}
		if msg.Type == MTMessage {
			messages <- msg.Data
		}
	}
}

func clientWriter(conn *websocket.Conn, ch <-chan string) {
	for msg := range ch {
		m := Message{
			Type: MTMessage,
			Data: msg,
		}
		if err := conn.WriteJSON(&m); err != nil {
			log.Println(err)
		}
	}
}
