package helpers

import (
	"dance-instructor/internal/config"
	"dance-instructor/internal/server"

	"github.com/labstack/echo/v4"
)

func NewServer(cfg *config.Config) *server.Server {
	return &server.Server{
		Echo:   echo.New(),
		Config: cfg,
	}
}
