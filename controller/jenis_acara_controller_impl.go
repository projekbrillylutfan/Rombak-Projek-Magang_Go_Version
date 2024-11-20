package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type JenisAcaraController struct {
	JenisAcaraService service.JenisAcaraService
	configuration.Config
}

func NewJenisAcaraController(jenisAcaraService *service.JenisAcaraService, config *configuration.Config) *JenisAcaraController {
	return &JenisAcaraController{
		JenisAcaraService: *jenisAcaraService,
		Config: *config,
	}
}

func(controller *JenisAcaraController)CreateJenisAcaraController(c *fiber.Ctx) error {
	var request *web.JenisAcaraCreateOrUpdate
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	controller.JenisAcaraService.CreateJenisAcaraService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code: 200,
		Message: "succes create jenis acara",
		Data: request,
	})
}