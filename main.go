package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/app"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	impl_controller "github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/controller/impl"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/exception"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/repository/impl_repo"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/route"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/service/impl_service"
)

func main() {
	config := configuration.New()
	database := app.NewDatabase(config)

	app.SeedAdminUser(database)

	// repository
	userRepository := impl_repo.NewUserRepository(database)
	bupatiRepository := impl_repo.NewBupatiRepository(database)
	lokasiRepository := impl_repo.NewLokasiRepository(database)
	jenisAcaraRepository := impl_repo.NewJenisAcaraRepository(database)
	agendaRepository := impl_repo.NewAgendaRepostiory(database)

	// service
	userService := impl_service.NewUserServiceImpl(&userRepository, &config)
	bupatiService := impl_service.NewBupatiServiceImpl(&bupatiRepository)
	lokasiService := impl_service.NewLokasiServiceImpl(&lokasiRepository)
	jenisAcaraService := impl_service.NewJenisAcaraServiceImpl(&jenisAcaraRepository)
	agendaService := impl_service.NewAgendaServiceImpl(&agendaRepository, &bupatiRepository, &jenisAcaraRepository, &lokasiRepository)

	// controller
	userController := impl_controller.NewUserControllerImpl(userService, config)
	bupatiController := impl_controller.NewBupatiControllerImpl(bupatiService, config)
	lokasiController := impl_controller.NewLokasiControllerImpl(lokasiService, config)
	jenisAcaraController := impl_controller.NewJenisAcaraControllerImpl(jenisAcaraService, config)
	agendaController := impl_controller.NewAgendaControllerImpl(agendaService, config)

	// fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	// route User dan Auth
	route.UserRouteAdmin(app, userController)
	route.UserRoute(app, userController)
	// route Bupati
	route.BupatiRouteAdmin(app, bupatiController)
	route.BupatiRoute(app, bupatiController)
	// route lokasi
	route.LokasiRouteAdmin(app, lokasiController)
	// route jenis acara
	route.JenisAcaraRouteAdmin(app, jenisAcaraController)
	// route agenda
	route.AgendaRouteAdmin(app, agendaController)

	err := app.Listen("localhost:3000")
	exception.PanicLogging(err)

}
