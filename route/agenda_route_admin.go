package route

import (
	"github.com/gofiber/fiber/v2"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/middleware"
)

func AgendaRouteAdmin (app *fiber.App, controller *impl_controller.AgendaController) {
	agendaRouteGroup := app.Group("/api/admin/agenda", middleware.AuthenticateJWT("ADMIN", controller.Config))

	// agenda create
	agendaRouteGroup.Post("/", controller.CreateAgendaController)
}