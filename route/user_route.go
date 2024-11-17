package route

import (
	"github.com/gofiber/fiber/v2"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
)

func UserRoute (app *fiber.App, controller *impl_controller.UserController) {
	userGroup := app.Group("/api/user")
	// user register
	userGroup.Post("/register", controller.RegisterUserController)
	
}