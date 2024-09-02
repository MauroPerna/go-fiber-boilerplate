package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Registra la solicitud entrante
		start := time.Now()
		log.Printf("Started %s %s for %s", c.Method(), c.OriginalURL(), c.IP())

		// Continuar con la solicitud
		err := c.Next()

		// Registra la respuesta saliente
		duration := time.Since(start)
		log.Printf("Completed %s %s in %v with status %d", c.Method(), c.OriginalURL(), duration, c.Response().StatusCode())

		return err
	}
}
