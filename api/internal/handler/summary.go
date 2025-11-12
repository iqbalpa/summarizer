package handler

import (
	"summarizer/internal/model"
	"summarizer/internal/service"

	"github.com/gofiber/fiber/v2"
)

// Router
func SummaryRouter(api fiber.Router, sh SummaryHandler) {
	api.Route("/summary", func(router fiber.Router) {
		router.Get("/:id", sh.GetSummary)
		router.Post("/", sh.CreateSummary)
		router.Get("/", sh.GetAllSummaries)
	})
}

// Handler
type SummaryHandler struct {
	ss service.ISummaryService
}

func NewSummaryHandler(ss service.ISummaryService) *SummaryHandler {
	return &SummaryHandler{
		ss: ss,
	}
}

func (sh *SummaryHandler) GetSummary(c *fiber.Ctx) error {
	id := c.Params("id")
	s, err := sh.ss.GetSummary(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(s)
}

func (sh *SummaryHandler) CreateSummary(c *fiber.Ctx) error {
	sr := new(model.SummaryRequest)
	if err := c.BodyParser(&sr); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	s, err := sh.ss.CreateSummary(sr.Title, sr.Content, "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusCreated).JSON(s)
}

func (sh *SummaryHandler) GetAllSummaries(c *fiber.Ctx) error {
	s, err := sh.ss.GetAllSummaries()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(s)
}
