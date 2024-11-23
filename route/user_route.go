package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
)

func UserRoute(app *fiber.App, controller controller.UserController) {
	userGroup := app.Group("/api/user")
	// user register
	userGroup.Post("/register", controller.RegisterUserController)
	// user login
	userGroup.Post("/login", controller.LoginUserController)
	// email verif
	userGroup.Get("/verify-email", controller.VerifyEmail)
	// forgot password
	userGroup.Post("/forgot-password", controller.ForgotPassword)
	// reset password
	userGroup.Post("/reset-password", controller.ResetPassword)

}
