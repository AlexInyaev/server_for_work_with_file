package main

import (
	"fmt"
	"github.com/AlexInyaev/server_for_work_with_file/internal/config"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Port)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello from echo")
	})

	httpPort := fmt.Sprintf(":%v", c.Port)
	e.Logger.Fatal(e.Start(httpPort))
}
