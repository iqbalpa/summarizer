package middleware

import (
	"strings"
	"summarizer/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func Authorization(c *fiber.Ctx) error {
	path := c.Path()
	if strings.Contains(path, "/auth/") {
		return c.Next()
	}

	authToken := c.Get("Authorization", "")
	if authToken == "" {
		return c.Status(fiber.StatusForbidden).JSON("unauthorized")
	}

	lst := strings.Split(authToken, " ")
	if len(lst) != 2 || lst[0] != "Bearer" {
		return c.Status(fiber.StatusForbidden).JSON("unauthorized")
	}

	claims, err := utils.ExtractClaims(lst[1])
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(err)
	}

	username, ok1 := claims["username"].(string)
	userId, ok2 := claims["userId"].(string)

	if !ok1 || !ok2 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "invalid token claims",
		})
	}

	c.Locals("username", username)
	c.Locals("userId", userId)

	return c.Next()
}
