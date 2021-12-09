package server

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"golculator/internal/server/numerical"
)

type ApiServer struct {
	Port string
}

func (s *ApiServer) RunServer() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(200, "Hello, World!")
	// })

	// e.POST("/", func(c echo.Context) error {
	// 	num := c.FormValue("number")
	// 	return c.String(200, num)
	// })

	e.POST("/", func(c echo.Context) error {
		num1, _ := numerical.ParseNumber(c.FormValue("number1"))
		num2, _ := numerical.ParseNumber(c.FormValue("number2"))
		operator := c.FormValue("operator")
		result := numerical.Calculate(num1, num2, operator)

		s := fmt.Sprintf("%f", result)
		return c.String(200, s)

	})
	e.Logger.Fatal(e.Start(":" + s.Port))
}
