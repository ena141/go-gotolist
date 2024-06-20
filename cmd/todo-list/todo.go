package main

import (
	"github.com/ena141/go-todolist/pkg/config"
	"github.com/ena141/go-todolist/pkg/routes"
)

func main() {
	config.InitDB()

	router := routes.InitRoutes()

	router.Run(":8080")
}
