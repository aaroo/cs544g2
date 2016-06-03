# Requirements Team2 MCGP

This document contains files that implement MCGP protocol. The following part will indicate each file’s function of realizing the protocol requirements in the paper: STATEFUL, CONCURRENT, SERVICE, CLIENT, UI.

CLIENT
client.go is the main client class. It provides a command like interface to the user, who can specify what message to send the server. I starts up and uses an ssl certificate to initialize and establish a  server connection

STATEFUL
Since there is a TCP connection between the client and server and there can be multiple back and forth messages on this connection, both client and server need to maintain state. So client.go is responsible for initiating the connection, participating in the handshake and version check and handling operations on the client side. Similarly, server.go is responsible for maintaining state on the server side. the function handleClient() specifically implements the DFA depending on what state it is in thereby what to respond to the client with.

UI
Since this is a client server protocol, like the state, the UI needs to be at both ends as well. Both sides in this implementation use a Go command line UI. On the client side, client.go provides the CLI for the user. Form the server side, main.go provides the cli to the user and starts up the server.

SERVICE
Not sure what service means here.

client.go : CLIENT, STATEFUL, UI  
main.go: CONCURRENT ,UI  
protocol.go: SERVICE  
server.go: STATEFUL, CONCURRENT, SERVICE, UI  
data.go: SERVICE  
