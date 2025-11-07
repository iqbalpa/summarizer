package handler

import (
	"summarizer/internal/model"
	"summarizer/internal/service"

	"github.com/gofiber/fiber/v2"
)

// Router
func UserRouter(api fiber.Router, uh UserHandler) {
	api.Route("/auth", func(router fiber.Router) {
		router.Get("/test", uh.TestUser)
		router.Post("/", uh.CreateUser)
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

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	us := new(model.UserRequest)
	if err := c.BodyParser(&us); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	u, err := uh.us.CreateUser(us.Name, us.Username, us.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusCreated).JSON(u)
}
