package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// StatusInList : checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// DBTransactionMiddleware : to setup the database transaction middleware
func DBTransactionMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		txHandle := db.Begin()
		log.Print("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Locals("db_trx", txHandle)
		c.Next()

		if StatusInList(c.Response().StatusCode(), []int{http.StatusOK, http.StatusCreated}) {
			log.Print("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trxxx commit error: ", err)
				return err
			}
		} else {
			log.Print("rolling back transaction due to status code: ", c.Response().StatusCode())
			txHandle.Rollback()
		}

		return nil
	}
}
