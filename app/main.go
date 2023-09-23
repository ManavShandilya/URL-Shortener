package main

import (
	"github.com/Garry028/url-shortener/database"
	"github.com/Garry028/url-shortener/routes"
)

func main() {
	database.Init()
	routes.SetupAndListen()
}
