package main

import (
	httpserver "yandexCourse/internal/httpServer"
)

func main() {
	server := httpserver.ServerInit()

	server.Run()
}
