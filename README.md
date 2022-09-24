# Go WebSocket Server and Client Sample

## WebSocket
WebSocket is a technology that enables bi-directional asynchronous communication between server and client.  
It is used to notify the client immediately or to push data from the server to the client.  

## Server Code
```Go
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
```

## Client Code
```Go
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

```

## Output Sample
### Server Side
11:20:43.126515 Receive : {"REST-Key":"REST-Value1"}, from Client  
11:20:43.126656 Send : {"REST-Key":"REST-Value1"}, to Client  
11:20:43.126666 Receive : {"REST-Key":"REST-Value2"}, from Client  
11:20:43.126674 Send : {"REST-Key":"REST-Value2"}, to Client  
11:20:43.126677 Receive : {"REST-Key":"REST-Value3"}, from Client  
11:20:43.126683 Send : {"REST-Key":"REST-Value3"}, to Client  
### Client Side
11:20:43.126473 Send : {"REST-Key":"REST-Value1"}, to Server  
11:20:43.126593 Send : {"REST-Key":"REST-Value2"}, to Server  
11:20:43.126603 Send : {"REST-Key":"REST-Value3"}, to Server  
11:20:43.126704 Receive : {"REST-Key":"REST-Value1"}, from Server  
11:20:43.126716 Receive : {"REST-Key":"REST-Value2"}, from Server  
11:20:43.126719 Receive : {"REST-Key":"REST-Value3"}, from Server 

## Output Image
![スクリーンショット 2022-09-24 11 20 56](https://user-images.githubusercontent.com/36861752/192076607-a64cb7c2-953e-4a6f-b569-09930e9627ca.png)

## Output Procedure
### Linux(or Mac)
#### Start server program.  
$ go build -o WS-Server WS-Server.go  
$ ./WS-Server  
#### Start client program.  
$ go build -o WS-Client WS-Client.go  
$ ./WS-Client 
### Windows
#### Start server program.  
$ go build -o WS-Server.exe WS-Server.go  
$ WS-Server.exe  
#### Start client program.  
$ go build -o WS-Client.exe WS-Client.go  
$ WS-Client.exe 
