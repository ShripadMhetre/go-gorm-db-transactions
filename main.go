package main

import (
	"log"

	"github.com/ShripadMhetre/go-gorm-db-transactions/controller"
	"github.com/ShripadMhetre/go-gorm-db-transactions/middleware"
	"github.com/ShripadMhetre/go-gorm-db-transactions/model"
	"github.com/ShripadMhetre/go-gorm-db-transactions/repository"
	"github.com/ShripadMhetre/go-gorm-db-transactions/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, _ := model.DBConnection()
	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	userService := service.NewUserService(userRepository)

	userController := controller.NewUserController(userService)

	users := app.Group("users")

	// User Endpoints
	users.Get("/", userController.GetAllUser)
	users.Post("/", userController.AddUser)

	// Money Transfer (Transaction) endpoint
	app.Post("/transfer-money", middleware.DBTransactionMiddleware(db), userController.TransferMoney)

	log.Fatal(app.Listen(":3000"))
}
