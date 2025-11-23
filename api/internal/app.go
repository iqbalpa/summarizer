package internal

import (
	"log"
	"summarizer/internal/handler"
	"summarizer/internal/repo"
	"summarizer/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func App() *fiber.App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := repo.ConnectDb()
	repo.MigrateDb(db)

	var userRepo repo.IUserRepository = repo.NewUserRepository(db)
	var userService service.IUserService = service.NewUserService(userRepo)
	var userHandler handler.UserHandler = *handler.NewUserHandler(userService)

	var summaryRepo repo.ISummaryRepository = repo.NewSummaryRepository(db)
	var summaryService service.ISummaryService = service.NewSummaryService(summaryRepo)
	var summaryHandler handler.SummaryHandler = *handler.NewSummaryHandler(summaryService)

	var jobRepo repo.IJobRepository = repo.NewJobRepository(db)
	var jobService service.IJobService = service.NewJobService(jobRepo, summaryService)
	var jobHandler handler.JobHandler = *handler.NewJobHandler(jobService)

	app := fiber.New()

	api := app.Group("/api/v1")
	handler.UserRouter(api, userHandler)
	handler.SummaryRouter(api, summaryHandler)
	handler.JobRouter(api, jobHandler)

	return app
}
