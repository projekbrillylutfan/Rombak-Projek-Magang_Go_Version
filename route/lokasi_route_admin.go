package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/middleware"
)

func LokasiRouteAdmin (app *fiber.App, controller controller.LokasiController) {
	lokasiGroupAdmin := app.Group("/api/admin/lokasi", middleware.AuthenticateJWT("ADMIN", controller.GetConfig()))
	// lokasi create
	lokasiGroupAdmin.Post("/", controller.CreateLokasiController)
	// lokasi find all
	lokasiGroupAdmin.Get("/", controller.FindAllLokasiController)
	// lokasi find by id
	lokasiGroupAdmin.Get("/:id", controller.FindByIdLokasiController)
	// lokasi update
	lokasiGroupAdmin.Put("/:id", controller.UpdateLokasiController)
	// lokasi delete
	lokasiGroupAdmin.Delete("/:id", controller.DeleteLokasiController)
}