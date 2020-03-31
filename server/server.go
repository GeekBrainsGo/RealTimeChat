package server

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/phadeyev/RealTimeChat/models"
)

type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader

	submutex  *sync.Mutex
	publisher *models.Publisher
}

func New() *Server {
	router := chi.NewRouter()
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 2014,
	}
	publisher := models.NewPublisher()
	ch := models.NewChannel()
	publisher.AddChannel("one", ch)
	serv := &Server{
		router:    router,
		upgrader:  upgrader,
		submutex:  &sync.Mutex{},
		publisher: publisher,
	}

	serv.bindRoutes()

	return serv
}

func (serv *Server) Start() error {
	return http.ListenAndServe(":8085", serv.router)
}

func SendPing(ws *websocket.Conn) {
	for {
		<-time.After(5 * time.Second)
		msg := models.Message{
			Type: models.MTPing,
		}
		if err := ws.WriteJSON(msg); err != nil {
			log.Printf("ws send ping err: %v", err)
			break
		}
	}
}
