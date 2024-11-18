package route

import (
	"github.com/gofiber/fiber/v2"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/middleware"
)

func UserRouteAdmin (app *fiber.App, controller *impl_controller.UserController) {
	userGroupAdmin := app.Group("/api/admin/user", middleware.AuthenticateJWT("ADMIN", controller.Config))
	// user create
	userGroupAdmin.Post("/", controller.CreateUserController)
	// user find by id
	userGroupAdmin.Get("/:id", controller.FindByIdUserController)
	// user find all
	userGroupAdmin.Get("/", controller.FindAllUserController)
	// user update
	userGroupAdmin.Put("/:id", controller.UpdateUserController)
	// user delete
	userGroupAdmin.Delete("/:id", controller.DeleteUserController)
}