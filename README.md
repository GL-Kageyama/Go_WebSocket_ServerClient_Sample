# Go WebSocket Server and Client Sample

## WebSocket

## Server Code
```Go

```

## Client Code
```Go

```

## Output Sample
### Server Side
$ go build -o WS-Server WS-Server.go  
$ ./WS-Server  
11:20:43.126515 Receive : {"REST-Key":"REST-Value1"}, from Client  
11:20:43.126656 Send : {"REST-Key":"REST-Value1"}, to Client  
11:20:43.126666 Receive : {"REST-Key":"REST-Value2"}, from Client  
11:20:43.126674 Send : {"REST-Key":"REST-Value2"}, to Client  
11:20:43.126677 Receive : {"REST-Key":"REST-Value3"}, from Client  
11:20:43.126683 Send : {"REST-Key":"REST-Value3"}, to Client  
### Client Side
$ go build -o WS-Client WS-Client.go  
$ ./WS-Client  
11:20:43.126473 Send : {"REST-Key":"REST-Value1"}, to Server  
11:20:43.126593 Send : {"REST-Key":"REST-Value2"}, to Server  
11:20:43.126603 Send : {"REST-Key":"REST-Value3"}, to Server  
11:20:43.126704 Receive : {"REST-Key":"REST-Value1"}, from Server  
11:20:43.126716 Receive : {"REST-Key":"REST-Value2"}, from Server  
11:20:43.126719 Receive : {"REST-Key":"REST-Value3"}, from Server 

## Output Image

## Output Procedure
1,  
2,  
