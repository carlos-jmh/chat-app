package chat

import (
	"time"

	"github.com/gorilla/websocket"
)

var (
	Rooms            = make(map[string]*Room)
	customTimeLayout = "01/02 03:04 PM"
)

type Room struct {
	Clients   map[*websocket.Conn]bool
	Messages  []string
	Broadcast chan string
}

func NewRoom(name string) *Room {
	room := &Room{
		Clients:   make(map[*websocket.Conn]bool),
		Messages:  make([]string, 0),
		Broadcast: make(chan string),
	}

	Rooms[name] = room

	return room
}

// Broadcaster stores and sends a message retrieved from the Broadcast channel to all clients
func (r *Room) Broadcaster() {
	for {
		msg := <-r.Broadcast

		// Add timestamp
		finalMsg := time.Now().Format(customTimeLayout) + " " + msg

		// Store
		r.Messages = append(r.Messages, finalMsg)

		// Send
		for conn := range r.Clients {
			err := conn.WriteMessage(websocket.TextMessage, []byte(finalMsg))
			if err != nil {
				conn.Close()
				delete(r.Clients, conn)
				continue
			}
		}
	}
}

// Restore restores the history of a chat room to a client
func (r *Room) Restore(conn *websocket.Conn) {
	for _, msg := range r.Messages {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			conn.Close()
			delete(r.Clients, conn)
			return
		}
	}
}
