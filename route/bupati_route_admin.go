package route

import (
	"github.com/gofiber/fiber/v2"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/middleware"
)

func BupatiRouteAdmin (app *fiber.App, controller *impl_controller.BupatiController) {
	bupatiGroupAdmin := app.Group("/api/admin/bupati", middleware.AuthenticateJWT("ADMIN", controller.Config))
	//bupati create
	bupatiGroupAdmin.Post("/", controller.CreateBupatiController)
	// bupati update
	bupatiGroupAdmin.Put("/:id", controller.UpdateBupatiController)
}