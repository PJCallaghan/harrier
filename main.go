package main

import (
	"harrier/internal/service"
)

func main() {
	server := service.NewServer()
	server.Start()
}
