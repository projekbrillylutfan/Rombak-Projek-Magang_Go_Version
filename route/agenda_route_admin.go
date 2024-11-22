package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/middleware"
)

func AgendaRouteAdmin (app *fiber.App, controller controller.AgendaController) {
	agendaRouteGroup := app.Group("/api/admin/agenda", middleware.AuthenticateJWT("ADMIN", controller.GetConfig()))

	// agenda create
	agendaRouteGroup.Post("/", controller.CreateAgendaController)
	// agenda find all
	agendaRouteGroup.Get("/", controller.FindAllAgendaController)
	// agenda find by id
	agendaRouteGroup.Get("/:id", controller.FindByIdAgendaController)
	// agenda update
	agendaRouteGroup.Put("/:id", controller.UpdateAgendaController)
}