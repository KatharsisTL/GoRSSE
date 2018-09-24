package SSE

import (
	"log"
	"net/http"
	"fmt"
)

type Server struct {
	Notifier chan []byte
	newClients chan chan []byte
	closingClients chan chan []byte
	clients map[chan []byte]bool
}

func NewServer() (server *Server) {
	server = &Server{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}
	go server.listen()
	return
}

func (server *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	messageChan := make(chan []byte)
	server.newClients <- messageChan
	defer func() {
		server.closingClients <- messageChan
	}()
	notify := rw.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		server.closingClients <- messageChan
	}()
	for {
		fmt.Fprintf(rw, "data: %s\n\n", <-messageChan)
		flusher.Flush()
	}
}

func (server *Server) listen() {
	for {
		select {
		case s := <-server.newClients:
			server.clients[s] = true
			log.Printf("Client added. %d registered clients", len(server.clients))
		case s := <-server.closingClients:
			delete(server.clients, s)
			log.Printf("Removed client. %d registered clients", len(server.clients))
		case event := <-server.Notifier:
			for clientMessageChan, _ := range server.clients {
				clientMessageChan <- event
			}
		}
	}
}

func (server *Server) Send(msg string) {
	server.Notifier <- []byte(msg)
}

func (server *Server) Listen(addr string) {
	http.ListenAndServe(addr, server)
}