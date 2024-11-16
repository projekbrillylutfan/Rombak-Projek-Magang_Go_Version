package impl_controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service/impl_service"
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
	userGroup.Get("/:id", controller.FindByIdUserController)
	userGroup.Get("/", controller.FindAllUserController)
	userGroup.Put("/:id", controller.UpdateUserController)
}

func (controller *UserController) CreateUserController(c *fiber.Ctx) error {
	var request *web.UserCreateOrUpdate
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	controller.UserService.CreateUserService(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success create user",
		Data: request,
	})
}

func (controller *UserController) FindByIdUserController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := impl_service.ConversionError(id)

	result := controller.UserService.FindByIdUserService(c.Context(), idInt64)

	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success get user",
		Data: result,
	})
}

func (controller *UserController)FindAllUserController(c *fiber.Ctx) error {
	result := controller.UserService.FindAllUserService(c.Context())
	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success get all user",
		Data: result,
	})
}

func (controller *UserController) UpdateUserController(c *fiber.Ctx) error {
	var request *web.UserCreateOrUpdate
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	idInt64 := impl_service.ConversionError(id)

	controller.UserService.UpdateUserService(c.Context(), request, idInt64)
	fmt.Println("isi payload di controller update ", request)
	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success update user",
		Data: request,
	})
}