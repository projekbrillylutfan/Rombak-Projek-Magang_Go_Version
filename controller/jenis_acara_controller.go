package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
)

type JenisAcaraController interface {
	CreateJenisAcaraController(c *fiber.Ctx) error
	FindAllJenisAcara(c *fiber.Ctx) error
	FindByIdJenisAcaraController(c *fiber.Ctx) error
	UpdateJenisAcaraController(c *fiber.Ctx) error
	DeleteJenisAcaraController(c *fiber.Ctx) error
	GetConfig() configuration.Config
}