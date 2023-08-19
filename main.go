package main

import (
	"thomasbaloyi/go-messanger/server"
)

func main() {
	listener := server.Listen("")
	if listener == nil {
		return
	}
}
