package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type LokasiControllerImpl struct {
	LokasiService service.LokasiService
	configuration.Config
}

func NewLokasiControllerImpl(lokasiService service.LokasiService, config configuration.Config) controller.LokasiController {
	return &LokasiControllerImpl{
		LokasiService: lokasiService,
		Config:        config,
	}
}

func (controller *LokasiControllerImpl) GetConfig() configuration.Config{
	return controller.Config
}

func (controller *LokasiControllerImpl) CreateLokasiController(c *fiber.Ctx) error {
	var request *web.LokasiCreateOrUpdate
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	controller.LokasiService.CreateLokasiService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success create lokasi",
		Data:    request,
	})
}

func (controller *LokasiControllerImpl) FindAllLokasiController(c *fiber.Ctx) error {
	result := controller.LokasiService.FindAllLokasiService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success get all lokasi",
		Data:    result,
	})
}

func (controller *LokasiControllerImpl) FindByIdLokasiController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)
	result := controller.LokasiService.FindByIdLokasiService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success get lokasi by id",
		Data:    result,
	})
}

func (controller *LokasiControllerImpl) UpdateLokasiController(c *fiber.Ctx) error {
	var request *web.LokasiCreateOrUpdate
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	controller.LokasiService.UpdateLokasiService(c.Context(), request, idInt64)
	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success update lokasi",
		Data:    request,
	})
}

func (controller *LokasiControllerImpl) DeleteLokasiController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	controller.LokasiService.DeleteLokasiService(c.Context(), idInt64)

	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code:    200,
		Message: "success delete lokasi",
		Data:    nil,
	})
}
