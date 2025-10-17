package main

import (
	"fiber-jwt-demo/internal/server"
	"fmt"
)

func main() {
	app := server.Setup()

	fmt.Println("Server is running on http://localhost:8080")
	app.Listen(":8080")
}