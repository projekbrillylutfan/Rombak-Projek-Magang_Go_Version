package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type BupatiController struct {
	BupatiService service.BupatiService
	configuration.Config
}

func NewBupatiController(bupatiService *service.BupatiService, config *configuration.Config) *BupatiController {
	return &BupatiController{
		BupatiService: *bupatiService,
		Config: *config,

	}
}

func (controller *BupatiController) CreateBupatiController(c *fiber.Ctx) error {
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