package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/middleware"
)

func JenisAcaraRouteAdmin (app *fiber.App, controller controller.JenisAcaraController) {
	jenisAcaraAdmin := app.Group("/api/admin/jenis-acara", middleware.AuthenticateJWT("ADMIN", controller.GetConfig()))

	jenisAcaraAdmin.Post("/", controller.CreateJenisAcaraController)
	jenisAcaraAdmin.Get("/", controller.FindAllJenisAcara)
	jenisAcaraAdmin.Get("/:id", controller.FindByIdJenisAcaraController)
	jenisAcaraAdmin.Put("/:id", controller.UpdateJenisAcaraController)
	jenisAcaraAdmin.Delete("/:id", controller.DeleteJenisAcaraController)
}