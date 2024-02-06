package main

import (
	"fmt"

	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"

	"socketexample/models"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
	Server *gosocketio.Server
)

func init() {
	Server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	fmt.Println("Socket Inititalize...")
}

func LoadSocket() {
	// socket connection
	Server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		fmt.Println("Connected", c.Id())
		c.Join("Room")
	})

	// socket disconnection
	Server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		fmt.Println("Disconnected", c.Id())

		// handles when someone closes the tab
		c.Leave("Room")
	})

	// chat socket
	Server.On("/chat", func(c *gosocketio.Channel, message models.Message) string {
		fmt.Println(message.Text)
		c.BroadcastTo("Room", "/message", message.Text)
		return "message sent successfully."
	})
}

func CreateRouter() {
	router = mux.NewRouter()
}

func main() {
	fmt.Println("main function")
	LoadSocket()
	CreateRouter()
}
