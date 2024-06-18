package infrastructure

import (
	"log/slog"
	"strconv"
	"time"

	router "github.com/yosuke7040/grpc-taskapp/backend/infrastructure/router"
)

type config struct {
	webServer     router.Server
	webServerPort router.Port
	ctxTimeout    time.Duration
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

func (c *config) WebServer() *config {
	s, err := router.NewWebServerFactory(
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		slog.Error("Error configured router server", err)
	}

	slog.Info("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		slog.Error("Error parsing port", err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
