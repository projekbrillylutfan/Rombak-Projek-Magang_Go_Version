package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: *userService,
	}
}

func (controller *UserController) Route (app *fiber.App) {
	userGroup := app.Group("/api/user")
	userGroup.Post("/", controller.CreateUserController)
}

func (controller *UserController) CreateUserController(c *fiber.Ctx) error {
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