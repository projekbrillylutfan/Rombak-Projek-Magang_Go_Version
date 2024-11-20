package route

import (
	"github.com/gofiber/fiber/v2"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/middleware"
)

func LokasiRouteAdmin (app *fiber.App, controller *impl_controller.LokasiController) {
	lokasiGroupAdmin := app.Group("/api/admin/lokasi", middleware.AuthenticateJWT("ADMIN", controller.Config))
	// lokasi create
	lokasiGroupAdmin.Post("/", controller.CreateLokasiController)
	// lokasi find all
	lokasiGroupAdmin.Get("/", controller.FindAllLokasiController)
	// lokasi find by id
	lokasiGroupAdmin.Get("/:id", controller.FindByIdLokasiController)
}