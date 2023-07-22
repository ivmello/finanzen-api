package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ivmello/finanzen-api/internal/adapter/api/controller"
)

func NewTransactionRouter(router fiber.Router, handler controller.TransactionControllerInterface) {
	group := router.Group("/transactions")
	group.Post("/", handler.CreateTransaction)
}
