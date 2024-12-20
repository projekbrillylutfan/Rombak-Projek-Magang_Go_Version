package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
)

func BupatiRoute (app *fiber.App, controller controller.BupatiController) {
	bupatGroup := app.Group("/api/bupati")

	// bupati find all
	bupatGroup.Get("/", controller.FindAllBupatiController)
	// bupati find by id
	bupatGroup.Get("/:id", controller.FindByIdBupatiController)
}