package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/middleware"
)

func BupatiRouteAdmin (app *fiber.App, controller controller.BupatiController) {
	bupatiGroupAdmin := app.Group("/api/admin/bupati", middleware.AuthenticateJWT("ADMIN", controller.GetConfig()))
	//bupati create
	bupatiGroupAdmin.Post("/", controller.CreateBupatiController)
	// bupati update
	bupatiGroupAdmin.Put("/:id", controller.UpdateBupatiController)
	// delete bupati
	bupatiGroupAdmin.Delete("/:id", controller.DeleteBupatiController)
}