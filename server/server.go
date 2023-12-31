package server

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func StartServer() net.Listener {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Listening on " + server.Addr().String())

	for {
		connection, err := server.Accept()

		if err != nil {
			fmt.Println("Error accepting to server:", err.Error())
			os.Exit(1)
		}

		fmt.Println("client connected")
		go processClientConnection(connection)
	}
}

func processClientConnection(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)

	if err != nil {
		fmt.Println("Error reading from connection:", err.Error())
	}

	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))

	if err != nil {
		fmt.Println("Error writing to connection:", err.Error())
	}

	connection.Close()
}
