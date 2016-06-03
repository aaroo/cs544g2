# Requirements Team2 MCGP

This document contains files that implement MCGP protocol. The following part will indicate each file’s function of realizing the protocol requirements in the paper: STATEFUL, CONCURRENT, SERVICE, CLIENT, UI.

```
client.go : CLIENT, STATEFUL
main.go: UI  
protocol.go: SERVICE  
server.go: STATEFUL, CONCURRENT, SERVICE
data.go: SERVICE
```

## CLIENT

`client.go` is the main client class. It provides a command line interface to the user, who can specify what message to send the server. I starts up and uses an ssl certificate to initialize and establish a  server connection

## STATEFUL

Since there is a TCP connection between the client and server, and there can be multiple back and forth messages on this connection, both client and server need to maintain state. So client.go is responsible for initiating the connection, participating in the handshake and version check and handling operations on the client side. Similarly, `server.go` is responsible for maintaining state on the server side. the function handleClient() specifically implements the DFA depending on what state it is in thereby what to respond to the client with.

## UI

Since this is a client server protocol, like the state, the UI needs to be at both ends as well. Both sides in this implementation use a command line UI. Commands for both client and the server are defined in `main.go`.

## SERVICE

Service is implemented in `server.go` which starts a listener that will concurrently accept as many connections as necessary on the predefined port (6666 by default),
