package application

import (
	"log"

	"dance-instructor/internal/config"
	"dance-instructor/internal/server"
	"dance-instructor/internal/server/routes"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
