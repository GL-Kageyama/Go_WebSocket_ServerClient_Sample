package main

import (
	"log"

	"golang.org/x/net/websocket"
)

// WebSocket Client Sample
func main() {
	log.SetFlags(log.Lmicroseconds)

	// WebSocket Dial
	ws, dialErr := websocket.Dial("ws://localhost:8000", "", "http://localhost:8000")
	if dialErr != nil {
		log.Fatal(dialErr)
	}
	defer ws.Close()

	// Send Message
	sendRestMsg(ws, `{"REST-Key":"REST-Value1"}`)
	sendRestMsg(ws, `{"REST-Key":"REST-Value2"}`)
	sendRestMsg(ws, `{"REST-Key":"REST-Value3"}`)

	// Receive Message Ligic
	var recvMsg string
	for {
		recvErr := websocket.Message.Receive(ws, &recvMsg)
		if recvErr != nil {
			log.Fatal(recvErr)
			break
		}
		log.Println("Receive : " + recvMsg + ", from Server")
	}
}

// Send Message Ligic
func sendRestMsg(ws *websocket.Conn, msg string) {
	sendErr := websocket.Message.Send(ws, msg)
	if sendErr != nil {
		log.Fatal(sendErr)
	}
	log.Println("Send : " + msg + ", to Server")
}
