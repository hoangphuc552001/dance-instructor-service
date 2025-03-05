package middlewares

import (
	"net/http"

	"dance-instructor/internal/server"

	"github.com/labstack/echo/v4"
)

type InternalMiddleware struct {
	server *server.Server
}

func NewInternalMiddleware(server *server.Server) *InternalMiddleware {
	return &InternalMiddleware{
		server: server,
	}
}

// InternalMiddleware is a middleware that will validate internal endpoint by secret key
func (i *InternalMiddleware) InternalAuthenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		secret := c.Request().Header.Get("Authorization")
		if secret == "" || secret != i.server.Config.Auth.InternalSecret {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				http.StatusText(http.StatusUnauthorized),
			)
		}
		return next(c)
	}
}
