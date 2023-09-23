package routes

import (
	"github.com/Garry028/url-shortener/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var router *fiber.App

func initHandlers() {
	router.Get("/goly", controller.GetAllRedirects)
	router.Get("/goly/:id", controller.GetGoly)
	router.Post("/goly", controller.CreateGoly)
	router.Patch("/goly", controller.UpdateGoly)
	router.Delete("/goly/:id", controller.DeleteGoly)
	router.Get("/r/:redirect", controller.Redirects)
}

func SetupAndListen() {
	router = fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	initHandlers()
	err := router.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
