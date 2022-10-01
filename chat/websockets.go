package chat

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func HandleNewClient(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	Rooms["main"].Clients[conn] = true

	// Restore chat room history
	Rooms["main"].Restore(conn)

	for {
		// Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		Rooms["main"].Broadcast <- string(msg)
	}
}
