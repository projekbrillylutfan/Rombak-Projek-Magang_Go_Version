package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/app"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository/impl_repo"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service/impl_service"
)



func main() {
    config := configuration.New()
    database := app.NewDatabase(config)

    // repository
    userRepository := impl_repo.NewUserRepository(database)

    // service
    userService := impl_service.NewUserServiceImpl(&userRepository)

    // controller
    userController := impl_controller.NewUserController(&userService)

    // fiber 
    app := fiber.New(configuration.NewFiberConfiguration())
    app.Use(recover.New())
    app.Use(cors.New())

    userController.Route(app)

    err := app.Listen("localhost:3000")
    exception.PanicLogging(err)

}