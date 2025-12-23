package main

import (
	"fmt"
	"github.com/AlexInyaev/server_for_work_with_file/internal/config"
	"log"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Port)

}
