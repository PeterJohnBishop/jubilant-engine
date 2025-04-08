package main

import (
	"fmt"
	"jupilant-engine/main.go/server"
)

func main() {
	fmt.Println("Jubilant-Engine, a Docker/Kubernetes playground.")
	server.StartServer()
}
