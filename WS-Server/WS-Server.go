package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// WebSocket Server Sample
func main() {
	log.SetFlags(log.Lmicroseconds)

	http.Handle(
		"/",
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			for {
				// Receive Logic
				var msg string
				recvErr := websocket.Message.Receive(ws, &msg)
				if recvErr != nil {
					log.Fatal(recvErr)
					break
				}
				log.Println("Receive : " + msg + ", from Client")

				// Send Logic
				sendErr := websocket.Message.Send(ws, msg)
				if sendErr != nil {
					log.Fatal(sendErr)
					break
				}
				log.Println("Send : " + msg + ", to Client")
			}
		}),
	)

	// WebSocket Listen and Serve
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
