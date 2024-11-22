package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type BupatiControllerImpl struct {
	BupatiService service.BupatiService
	configuration.Config
}

func NewBupatiControllerImpl(bupatiService service.BupatiService, config configuration.Config) controller.BupatiController {
	return &BupatiControllerImpl{
		BupatiService: bupatiService,
		Config: config,

	}
}

func (controller *BupatiControllerImpl) GetConfig() configuration.Config{
	return controller.Config
}

func (controller *BupatiControllerImpl) CreateBupatiController(c *fiber.Ctx) error {
	var request *web.BupatiCreateOrUpdate
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	controller.BupatiService.CreateBupatiService(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success create bupati",
		Data: request,
	})
}

func (controller *BupatiControllerImpl) FindAllBupatiController(c *fiber.Ctx) error {
	result := controller.BupatiService.FindAllBupatiService(c.Context())
	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success get all bupati",
		Data: result,
	})
}

func (controller *BupatiControllerImpl) FindByIdBupatiController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	result := controller.BupatiService.FindByIdService(c.Context(), idInt64)

	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success get bupati by id",
		Data: result,
	})
}

func (controller *BupatiControllerImpl) UpdateBupatiController(c *fiber.Ctx) error {
	var request *web.BupatiCreateOrUpdate
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	controller.BupatiService.UpdateBupatiService(c.Context(), request, idInt64)
	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success update bupati",
		Data: request,
	})
}

func (controller *BupatiControllerImpl) DeleteBupatiController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	controller.BupatiService.DeleteBupatiService(c.Context(), idInt64)

	return c.Status(fiber.StatusOK).JSON(&web.GeneralResponse{
		Code: 200,
		Message: "success delete bupati",
		Data: nil,
	})
}