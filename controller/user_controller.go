package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	CreateUserController(c *fiber.Ctx) error
}