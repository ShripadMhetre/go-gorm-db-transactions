package controller

import (
	"log"
	"net/http"

	"github.com/ShripadMhetre/go-gorm-db-transactions/model"
	"github.com/ShripadMhetre/go-gorm-db-transactions/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserController : represent the user's controller contract
type UserController interface {
	AddUser(*fiber.Ctx)
	GetAllUser(*fiber.Ctx)
	TransferMoney(*fiber.Ctx)
}

type userController struct {
	userService service.UserService
}

// NewUserController -> returns new user controller
func NewUserController(s service.UserService) UserController {
	return userController{
		userService: s,
	}
}

func (u userController) GetAllUser(c *fiber.Ctx) {
	log.Print("[UserController]...get all Users")

	users, err := u.userService.GetAll()
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(err.Error())
		return
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"data": users})
}

func (u userController) AddUser(c *fiber.Ctx) {
	log.Print("[UserController]...add User")
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	user, err := u.userService.Save(user)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Error while saving user"})
		return
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"data": user})
}

func (u userController) TransferMoney(c *fiber.Ctx) {
	log.Print("[UserController]...get all Users")

	txHandle := c.Locals("db_trx").(*gorm.DB)

	var moneyTransfer model.MoneyTransfer
	if err := c.BodyParser(&moneyTransfer); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return
	}

	if err := u.userService.WithTrx(txHandle).IncrementMoney(moneyTransfer.Receiver, moneyTransfer.Amount); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Error while incrementing money"})
		txHandle.Rollback()
		return
	}

	if err := u.userService.WithTrx(txHandle).DecrementMoney(moneyTransfer.Giver, moneyTransfer.Amount); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Error while decrementing money"})
		txHandle.Rollback()
		return
	}

	if err := txHandle.Commit().Error; err != nil {
		log.Print("trx commit error: ", err)
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"msg": "Successfully Money Transferred"})
}
