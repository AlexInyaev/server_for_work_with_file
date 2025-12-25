package main

import (
	"fmt"
	"github.com/AlexInyaev/server_for_work_with_file/internal/config"
	"github.com/AlexInyaev/server_for_work_with_file/internal/executor"
	"github.com/AlexInyaev/server_for_work_with_file/internal/handler"
	"github.com/AlexInyaev/server_for_work_with_file/internal/infrastructure/repository"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	//config
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Port)
	//infrastructure
	repo := repository.New()
	//executors
	writeText := executor.NewWriteText(repo)
	//routes
	e := echo.New()
	h := handler.NewHandler(e, writeText)

	h.InitRoutes()
	//run
	httpPort := fmt.Sprintf(":%v", c.Port)
	e.Logger.Fatal(e.Start(httpPort))
}
