package impl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) controller.UserController {
	return &UserControllerImpl{
		UserService: *userService,
	}
}

func (controller *UserControllerImpl) CreateUserController(c *fiber.Ctx) error {
	var request *web.UserCreate
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	controller.UserService.CreateUserService(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success create user",
		Data: request,
	})
}