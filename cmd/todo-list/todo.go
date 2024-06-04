package main

import (
	"github.com/ena141/go-gotolist/pkg/config"
	"github.com/ena141/go-gotolist/pkg/routes"
)

func main() {
	config.InitDB()

	router := routes.InitRoutes()

	router.Run(":8080")
}
