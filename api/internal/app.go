package internal

import (
	"summarizer/internal/handler"
	"summarizer/internal/repo"
	"summarizer/internal/service"

	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	db := repo.ConnectDb()
	repo.MigrateDb(db)
	
	var userRepo repo.IUserRepository = repo.NewUserRepository(db)
	var userService service.IUserService = service.NewUserService(userRepo)
	var userHandler handler.UserHandler = *handler.NewUserHandler(userService)

	app := fiber.New()
	api := app.Group("/api/v1")
	handler.UserRouter(api, userHandler)

	return app
}
