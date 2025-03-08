package routes

import (
	"dance-instructor/internal/server"
	"dance-instructor/internal/server/handlers"
	"dance-instructor/internal/server/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *server.Server) {
	// Init handlers
	instructionHandler := handlers.NewInstructorHandler(server)

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	internalMiddleware := middlewares.NewInternalMiddleware(server)
	v1 := server.Echo.Group("/api/v1")
	setupInstructorRoutes(v1, internalMiddleware, instructionHandler)
}

func setupInstructorRoutes(v1 *echo.Group, internalMiddleware *middlewares.InternalMiddleware, instructionHandler *handlers.InstructorHandlers) {
	instructorGroup := v1.Group("/instructors", internalMiddleware.InternalAuthenticate)

	instructorGroup.GET("", instructionHandler.GetAllInstructors)
	instructorGroup.GET("/:id", instructionHandler.GetInstructorByID)
	instructorGroup.POST("", instructionHandler.PostInstructor)
	instructorGroup.PUT("/:id", instructionHandler.PutInstructor)
	instructorGroup.DELETE("/:id", instructionHandler.DeleteInstructor)
}
