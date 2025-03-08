package responses

import (
	"github.com/labstack/echo/v4"
)

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(c echo.Context, statusCode int, data any) error {
	//nolint // context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//nolint // context.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	//nolint // context.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization")
	return c.JSON(statusCode, data)
}

func MessageResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Data{
		Code:    statusCode,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return Response(c, statusCode, Error{
		Code:  statusCode,
		Error: message,
	})
}

type ResponseDataForGetList struct {
	Total int64 `json:"total"`
	List  any   `json:"list"`
}

type ResponseData struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
