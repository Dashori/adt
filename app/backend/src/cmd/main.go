package main

import (
	server "app/internal/server"
	"os"
)

func main() {
	host := os.Getenv("BACKEND_HOST")
	port := os.Getenv("BACKEND_PORT")

	server.SetupServer().Run(host + ":" + port)
}
