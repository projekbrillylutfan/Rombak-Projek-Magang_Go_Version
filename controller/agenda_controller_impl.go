package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type AgendaController struct {
	AgendaService service.AgendaService
	configuration.Config
}

func NewAgendaController(agendaService *service.AgendaService, config *configuration.Config) *AgendaController {
	return &AgendaController{
		AgendaService: *agendaService,
		Config:        *config,
	}
}

func (controller *AgendaController)CreateAgendaController(c *fiber.Ctx) error {
	var request *web.AgendaCreateOrUpdate
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	controller.AgendaService.CreateAgendaService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code: 200,
		Message: "success create agenda",
		Data: request,
	})
}

func (controller *AgendaController) FindAllAgendaController(c *fiber.Ctx) error {
	result := controller.AgendaService.FindAllAgendaService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code: 200,
		Message: "success get all agenda",
		Data: result,
	})
}

func (controller *AgendaController)FindByIdAgendaController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	result := controller.AgendaService.FindByIdAgendaService(c.Context(), idInt64)

	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code: 200,
		Message: "succes get agenda by id",
		Data: result,
	})
}
