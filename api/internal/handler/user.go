package handler

import (
	"summarizer/internal/service"

	"github.com/gofiber/fiber/v2"
)

// Router
func UserRouter(api fiber.Router, uh UserHandler) {
	api.Route("/auth", func(router fiber.Router) {
		router.Get("/test", uh.TestUser)
		router.Post("/", uh.TestUser)
	})
}

// Handler
type UserHandler struct {
	us service.IUserService
}

func NewUserHandler(us service.IUserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func (uh *UserHandler) TestUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"message": "hello",
	})
}
