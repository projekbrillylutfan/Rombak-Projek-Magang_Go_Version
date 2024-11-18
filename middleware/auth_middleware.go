package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/configuration"
	"github.com/projekbrillylutfan/Rombak-Projek-Magang_Go_Version/model/web"
)

func AuthenticateJWT(role string, config configuration.Config) func(*fiber.Ctx) error {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return func(ctx *fiber.Ctx) error {
		// Ambil token dari Authorization header
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.
				Status(fiber.StatusBadRequest).
				JSON(&web.GeneralResponse{
					Code:    400,
					Message: "Bad Request",
					Data:    "Missing Authorization header",
				})
		}

		// Pastikan format header adalah "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return ctx.
				Status(fiber.StatusBadRequest).
				JSON(&web.GeneralResponse{
					Code:    400,
					Message: "Bad Request",
					Data:    "Invalid Authorization header format",
				})
		}

		tokenString := parts[1]

		// Parse dan validasi token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Validasi algoritma
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(&web.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid or expired JWT",
				})
		}

		// Ambil claims dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(&web.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid JWT claims",
				})
		}

		// Ambil roles dari claims
		roles, ok := claims["roles"].(string)
		if !ok {
			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(web.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid roles in JWT",
				})
		}

		// Log dan validasi role
		configuration.NewLogger().Info("role function ", role, " role user ", roles)

		// Cek apakah role cocok dengan yang dibutuhkan
		if roles == role {
			return ctx.Next() // Jika cocok, lanjutkan ke handler berikutnya
		}

		return ctx.
			Status(fiber.StatusUnauthorized).
			JSON(&web.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    "Invalid Role",
			})
	}
}
