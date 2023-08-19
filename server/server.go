package server

import (
	"fmt"
	"net"
)

func ListenAndServe(port string) {

	listener, err := net.Listen("tcp4", "localhost:"+port)

	if err != nil {
		fmt.Println("Could not start server: ", err)
		return
	} else {
		fmt.Println("Listening and serving on the address:", listener.Addr().String())
	}

	connection, err := listener.Accept()

	if err != nil {
		fmt.Println("Could not accept/open connection: ", err)
	} else {
		fmt.Println("Connection established: ", connection)
	}

	err = connection.Close()

	if err != nil {
		fmt.Println("Could not close connection: ", err)
	} else {
		fmt.Println("Connection closed: ", connection)
	}

}
