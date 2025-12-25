package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"github.com/AlexInyaev/server_for_work_with_file/internal/config"
	"github.com/AlexInyaev/server_for_work_with_file/internal/executor"
	"github.com/AlexInyaev/server_for_work_with_file/internal/handler"
	"github.com/AlexInyaev/server_for_work_with_file/internal/infrastructure/repository"
)

func main() {
	// config
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// infrastructure
	repo := repository.New()
	// executors
	exWriteText := executor.NewWriteText(repo)
	// routes
	e := echo.New()
	h := handler.NewHandler(e, exWriteText)

	h.InitRoutes()
	// run
	httpPort := fmt.Sprintf(":%v", c.Port)
	e.Logger.Fatal(e.Start(httpPort))
}
