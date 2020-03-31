package server

import (
	eventchannel "chat/event-channel"
	"chat/message"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// ApplyHandlers applies all server handlers.
func (s *Server) ApplyHandlers() {
	s.router.Handle("/*", http.FileServer(http.Dir("./web")))
	s.router.Get("/socket", s.socketHandler)
}

func (s *Server) socketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	defer ws.Close()
	if err != nil {
		log.Println(err)
		return
	}

	user := eventchannel.NewUser(uuid.New().String(), ws)
	ch := eventchannel.NewChannel()
	ch.Subscribe(user)

	s.mux.Lock()
	s.pub.AddChannel(user.Name, ch)
	s.mux.Unlock()

	defer func() {
		s.mux.Lock()
		if err := s.pub.DeleteChannel(user.Name); err != nil {
			log.Println(err)
		}
		s.mux.Unlock()
	}()

	for {
		msg := message.Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			log.Println(err)
			return
		}
		if msg.Type == message.MTPong {
			continue
		}
		if msg.Type == message.MTMessage {
			s.mux.Lock()
			s.pub.Send(msg.Data, msg.Channels...)
			s.mux.Unlock()
		}
	}
}
