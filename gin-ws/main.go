package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebsocket(c *gin.Context) {
	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade:", err.Error())
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read messag error:", err.Error())
			break
		}
		fmt.Printf("recieve message: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("write message error:", err.Error())
			break
		}
	}
}

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"message": "web server power by gin",
		})
	})

	engine.GET("/ws", handleWebsocket)

	engine.Run(":8080")
}
