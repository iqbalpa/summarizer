package handler

import (
	"summarizer/internal/middleware"
	"summarizer/internal/model"
	"summarizer/internal/service"

	"github.com/gofiber/fiber/v2"
)

// Router
func JobRouter(api fiber.Router, jh JobHandler) {
	api.Use(middleware.Authorization)
	api.Route("job", func(router fiber.Router) {
		router.Get("/:id", jh.GetJob)
		router.Post("/summarize", jh.CreateJob)
	})
}

// Handler
type JobHandler struct {
	js service.IJobService
}

func NewJobHandler(js service.IJobService) *JobHandler {
	return &JobHandler{
		js: js,
	}
}

func (jh *JobHandler) GetJob(c *fiber.Ctx) error {
	id := c.Params("id")
	j, err := jh.js.GetJob(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(j)
}

func (jh *JobHandler) CreateJob(c *fiber.Ctx) error {
	jr := new(model.JobRequest)
	if err := c.BodyParser(&jr); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	userId := c.Locals("userId").(string)
	j, err := jh.js.CreateJob(jr.Title, jr.Content, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(j)
}
