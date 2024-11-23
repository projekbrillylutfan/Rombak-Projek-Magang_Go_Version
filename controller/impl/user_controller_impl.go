package impl_controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service"
)

type UserControllerImpl struct {
	service.UserService
	configuration.Config
}

func NewUserControllerImpl(userService service.UserService, config configuration.Config) controller.UserController {
	return &UserControllerImpl{
		UserService: userService,
		Config:      config,
	}
}

func (controller *UserControllerImpl) GetConfig() configuration.Config{
	return controller.Config
}

func (controller *UserControllerImpl) CreateUserController(c *fiber.Ctx) error {
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

func (controller *UserControllerImpl) FindByIdUserController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	result := controller.UserService.FindByIdUserService(c.Context(), idInt64)

	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success get user",
		Data: result,
	})
}

func (controller *UserControllerImpl)FindAllUserController(c *fiber.Ctx) error {
	result := controller.UserService.FindAllUserService(c.Context())
	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success get all user",
		Data: result,
	})
}

func (controller *UserControllerImpl) UpdateUserController(c *fiber.Ctx) error {
	var request *web.UserCreateOrUpdate
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	controller.UserService.UpdateUserService(c.Context(), request, idInt64)
	fmt.Println("isi payload di controller update ", request)
	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success update user",
		Data: request,
	})
}

func (controller *UserControllerImpl) DeleteUserController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	controller.UserService.DeleteUserService(c.Context(), idInt64)

	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success delete user",
		Data: nil,
	})
}

func (controller *UserControllerImpl)RegisterUserController(c *fiber.Ctx) error {
	var request *web.UserRegister
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.RegisterUserService(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success register user, Check your email for verification.",
		Data: result,
	})
}

func (controller *UserControllerImpl)LoginUserController(c *fiber.Ctx) error {
	var request *web.UserLogin
	err := c.BodyParser(&request)
	exception.PanicLogging(err)


	result := controller.UserService.Authentication(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(web.GeneralResponse{
		Code: 200,
		Message: "success login user",
		Data: result,
	})
}

func (controller *UserControllerImpl)VerifyEmail(ctx *fiber.Ctx) error {
	token := ctx.Query("token")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Token is required")
	}

	err := controller.UserService.VerifyEmailService(ctx.Context(), token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"message": "Email successfully verified!",
	})
}

func (controller *UserControllerImpl)ForgotPassword(ctx *fiber.Ctx) error {
	var request *web.UserEmail

	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := controller.UserService.ForgotPasswordService(ctx.Context(), request.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"message": "Password reset link sent to your email!",
	})
}

func (controller *UserControllerImpl)ResetPassword(ctx *fiber.Ctx) error {
	var request *web.RestPassUser

	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := controller.UserService.ResetPasswordService(ctx.Context(), request.Token, request.NewPassword)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"message": "Password has been reset successfully!",
	})
}