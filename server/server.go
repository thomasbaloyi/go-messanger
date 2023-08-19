package server

import (
	"fmt"
	"net"
)

func Listen(port string) net.Listener {

	listener, err := net.Listen("tcp4", "localhost:"+port)

	if err != nil {
		fmt.Println("Could not start server: ", err)
		return nil
	} else {
		fmt.Println("Listening and serving on the address:", listener.Addr().String())
	}

	return listener
}

func Connect(listener net.Listener) (net.Conn, error) {
	return listener.Accept()
}

func StopListening(connection net.Conn) error {
	return connection.Close()
}
