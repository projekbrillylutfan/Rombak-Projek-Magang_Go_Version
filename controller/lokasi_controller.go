package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
)

type LokasiController interface {
	CreateLokasiController(c *fiber.Ctx) error
	FindAllLokasiController(c *fiber.Ctx) error
	FindByIdLokasiController(c *fiber.Ctx) error
	UpdateLokasiController(c *fiber.Ctx) error
	DeleteLokasiController(c *fiber.Ctx) error
	GetConfig() configuration.Config
}
