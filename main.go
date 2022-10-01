package main

import (
	"chat-app/chat"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", chat.HomePage)

	r.GET("/ws", chat.HandleNewClient)

	main := chat.NewRoom("main")
	go main.Broadcaster()

	r.Run(":5050")
}
