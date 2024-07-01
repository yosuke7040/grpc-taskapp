package main

import (
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure"
)

func main() {
	app := infrastructure.NewConfig()

	app.WebServerPort("8080").
		WebServer().
		Start()
}
