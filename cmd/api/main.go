// @title Swagger Example API (net/http)
// @version 1.0
// @description This is a sample server using Go's default net/http
// @host localhost:8080
// @BasePath /

package main

import (
	"github.com/Arash-mlk24/simple-task-manager-web-backend/internal/server"
)

func main() {
	server.Config()
}
