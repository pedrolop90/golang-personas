package main

import (
	"primero_test/server"
)

func main() {
	server := server.NewFabricaServer().ServerFactory()
	server.Init()
}
