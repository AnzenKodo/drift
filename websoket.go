package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
)

type Relay struct {
	connections map[*websocket.Conn]bool
	mutex       sync.Mutex
}

func (r *Relay) addConnection(conn *websocket.Conn) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.connections[conn] = true
}

func (r *Relay) removeConnection(conn *websocket.Conn) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.connections, conn)
}

func (r *Relay) broadcast(message []byte) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for conn := range r.connections {
		err := conn.Write(context.Background(), websocket.MessageText, message)
		if err != nil {
			log.Printf("error broadcasting to connection: %v", err)
			conn.CloseNow()
			delete(r.connections, conn)
		}
	}
}

func (r *Relay) handleConnection(conn *websocket.Conn) {
	defer func() {
		r.removeConnection(conn)
		conn.CloseNow()
	}()

	for {
		_, reader, err := conn.Reader(context.Background())
		if err != nil {
			log.Printf("read error: %v", err)
			break
		}

		message := make([]byte, 512)
		n, err := reader.Read(message)
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}

		message = message[:n]
		fmt.Printf("Received: %s\n", message)

		// Basic Nostr message handling
		switch {
		case string(message[:6]) == "EVENT ":
			// Broadcast the event to all connected clients
			r.broadcast(message)
		case string(message[:4]) == "REQ ":
			// Handle subscription requests
			// You might add logic here to filter events based on the REQ message
			conn.Write(context.Background(), websocket.MessageText, []byte("Acknowledged REQ"))
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request, ip string, ua string) {
    conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
    if err != nil {
		log.Println(err)
		return
	}
    log.Println("[Engine WS]", ip, " connected ", ua)
    conn.CloseNow()

    relay := &Relay{
		connections: make(map[*websocket.Conn]bool),
	}
	relay.addConnection(conn)
	go relay.handleConnection(conn)
}
