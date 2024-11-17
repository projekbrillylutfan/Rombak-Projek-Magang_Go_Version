package route

import (
	"github.com/gofiber/fiber/v2"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
)

func UserRoute (app *fiber.App, controller *impl_controller.UserController) {
	userGroup := app.Group("/api/user")
	// user create
	userGroup.Post("/", controller.CreateUserController)
	// user find by id
	userGroup.Get("/:id", controller.FindByIdUserController)
	// user find all
	userGroup.Get("/", controller.FindAllUserController)
	// user update
	userGroup.Put("/:id", controller.UpdateUserController)
	// user delete
	userGroup.Delete("/:id", controller.DeleteUserController)
	// register user by verif
	userGroup.Post("/register", controller.RegisterUserController)
}