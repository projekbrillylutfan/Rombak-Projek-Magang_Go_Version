package exception

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(web.GeneralResponse{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(web.GeneralResponse{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(web.GeneralResponse{
			Code:    401,
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(web.GeneralResponse{
		Code:    500,
		Message: "General Error",
		Data:    err.Error(),
	})
}