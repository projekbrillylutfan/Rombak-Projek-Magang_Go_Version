package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type LokasiController struct {
	LokasiService service.LokasiService
	configuration.Config
}

func NewLokasiController(lokasiService *service.LokasiService, config *configuration.Config) *LokasiController {
	return &LokasiController{
		LokasiService: *lokasiService,
		Config: *config,
	}
}

func (controller *LokasiController) CreateLokasiController(c *fiber.Ctx) error {
	var request *web.LokasiCreateOrUpdate
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	controller.LokasiService.CreateLokasiService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code: 200,
		Message: "success create lokasi",
		Data: request,
	})
}

func (controller *LokasiController) FindAllLokasiController(c *fiber.Ctx) error {
	result:= controller.LokasiService.FindAllLokasiService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code: 200,
		Message: "success get all lokasi",
		Data: result,
	})
}