package middleware

import (
	"github.com/MauroPerna/go-fiber-backend-test/services/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(authService *auth.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener el token de la cabecera de autorización
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization token not provided",
			})
		}

		// Eliminar el prefijo "Bearer " si existe
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		// Verificar el token
		decodedToken, err := authService.VerifyAccessToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Guardar los detalles del usuario en el contexto
		c.Locals("user", decodedToken)

		// Continuar con la siguiente handler
		return c.Next()
	}
}
