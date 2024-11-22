package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type JenisAcaraControllerImpl struct {
	JenisAcaraService service.JenisAcaraService
	configuration.Config
}

func NewJenisAcaraControllerImpl(jenisAcaraService service.JenisAcaraService, config configuration.Config) controller.JenisAcaraController {
	return &JenisAcaraControllerImpl{
		JenisAcaraService: jenisAcaraService,
		Config:            config,
	}
}

func (controller *JenisAcaraControllerImpl) GetConfig() configuration.Config{
	return controller.Config
}

func (controller *JenisAcaraControllerImpl) CreateJenisAcaraController(c *fiber.Ctx) error {
	var request *web.JenisAcaraCreateOrUpdate
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	controller.JenisAcaraService.CreateJenisAcaraService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "succes create jenis acara",
		Data:    request,
	})
}

func (controller *JenisAcaraControllerImpl) FindAllJenisAcara(c *fiber.Ctx) error {
	result := controller.JenisAcaraService.FindAllJenisAcaraService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success get all jenis acara",
		Data:    result,
	})
}

func (controller *JenisAcaraControllerImpl) FindByIdJenisAcaraController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	result := controller.JenisAcaraService.FindByIdJenisAcaraService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success get jenis acara by id",
		Data:    result,
	})
}

func (controller *JenisAcaraControllerImpl) UpdateJenisAcaraController(c *fiber.Ctx) error {
	var request *web.JenisAcaraCreateOrUpdate
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	controller.JenisAcaraService.UpdateJenisAcaraService(c.Context(), request, idInt64)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success update jenis acara",
		Data:    request,
	})
}

func (controller *JenisAcaraControllerImpl) DeleteJenisAcaraController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	controller.JenisAcaraService.DeleteJenisAcaraService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success delete jenis acara",
		Data:    nil,
	})
}
