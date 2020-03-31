package server

import (
	eventchannel "chat/event-channel"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

/* type client eventchannel.Subscriber

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan Message) // all incoming client messages
)
*/

// Server stands for server struct.
type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader
	pub      *eventchannel.Publisher
	mux      *sync.Mutex
}

// New creates new server.
func New() *Server {
	router := chi.NewRouter()

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	s := &Server{
		router:   router,
		upgrader: upgrader,
		pub:      eventchannel.NewPublisher(),
		mux:      &sync.Mutex{},
	}

	s.ApplyHandlers()
	return s
}

// Start starts server.
func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.router)
}
