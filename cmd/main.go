package main

import (
	"fmt"

	application "dance-instructor"
	"dance-instructor/docs"
	"dance-instructor/internal/config"
)

//	@title			Dance Instructor App
//	@version		2.0
//	@description	This is a Dance Instructor app.

// @BasePath	/
func main() {
	cfg := config.NewConfig()
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	application.Start(cfg)
}
