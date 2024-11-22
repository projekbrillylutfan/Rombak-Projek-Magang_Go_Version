package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
)

type BupatiController interface {
	CreateBupatiController(c *fiber.Ctx) error
	FindAllBupatiController(c *fiber.Ctx) error
	FindByIdBupatiController(c *fiber.Ctx) error
	UpdateBupatiController(c *fiber.Ctx) error
	DeleteBupatiController(c *fiber.Ctx) error
	GetConfig() configuration.Config
}