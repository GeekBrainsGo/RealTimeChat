package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

type client struct {
	out  chan<- string // an outgoing message channel
	name string        // uuid?
}

var (
	entering = make(chan *client)
	leaving  = make(chan *client)
	messages = make(chan string) // all incoming client messages
)

// Server stands for server struct.
type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader
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
	}

	s.ApplyHandlers()
	return s
}

// Start starts server.
func (s *Server) Start() error {
	go broadcaster()
	return http.ListenAndServe(":8080", s.router)
}

func broadcaster() {
	clients := make(map[*client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				select {
				case cli.out <- msg:
				default:
				}
			}
		case ent := <-entering:
			clients[ent] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.out)
		}
	}
}
