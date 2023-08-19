package main

import (
	"thomasbaloyi/go-messanger/client"
	"thomasbaloyi/go-messanger/server"
)

func main() {
	go server.StartServer()

	client.ConnectToServer()
}
