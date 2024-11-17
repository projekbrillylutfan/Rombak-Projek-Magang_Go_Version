package route

import (
	"github.com/gofiber/fiber/v2"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
)

func UserRoute (app *fiber.App, controller *impl_controller.UserController) {
	userGroup := app.Group("/api/user")
	userGroup.Post("/", controller.CreateUserController)
	userGroup.Get("/:id", controller.FindByIdUserController)
	userGroup.Get("/", controller.FindAllUserController)
	userGroup.Put("/:id", controller.UpdateUserController)
	userGroup.Delete("/:id", controller.DeleteUserController)
}