package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
)

type AgendaController interface {
	CreateAgendaController(c *fiber.Ctx) error
	FindAllAgendaController(c *fiber.Ctx) error
	FindByIdAgendaController(c *fiber.Ctx) error
	UpdateAgendaController(c *fiber.Ctx) error
	GetConfig() configuration.Config
	DeleteAgendaController(c *fiber.Ctx) error
}