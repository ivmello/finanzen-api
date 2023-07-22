package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ivmello/finanzen-api/internal/adapter/api/exception"
	"github.com/ivmello/finanzen-api/internal/adapter/api/response"
	"github.com/ivmello/finanzen-api/internal/adapter/api/validation"
	"github.com/ivmello/finanzen-api/internal/application/transaction/dto"
	"github.com/ivmello/finanzen-api/internal/application/transaction/usecase"
)

type TransactionControllerInterface interface {
	CreateTransaction(ctx *fiber.Ctx) error
}

type controller struct {
	createTransactionUsecase usecase.CreateTransactionUseCaseInterface
}

func NewTransactionController(createTransactionUsecase usecase.CreateTransactionUseCaseInterface) TransactionControllerInterface {
	return &controller{
		createTransactionUsecase,
	}
}

func (h *controller) CreateTransaction(ctx *fiber.Ctx) error {
	var input dto.CreateTransactionInputDto

	errParser := ctx.BodyParser(&input)
	if errParser != nil {
		errorResponse := exception.NewBadGatewayError(errParser.Error())
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	errValidation := validation.ValidateStruct(input)
	if errValidation != nil {
		return ctx.Status(errValidation.Code).JSON(errValidation)
	}

	output, errUsecase := h.createTransactionUsecase.Execute(input)

	if errUsecase != nil {
		errorResponse := exception.NewBadGatewayError(errUsecase.Error())
		return ctx.JSON(errorResponse)
	}

	response := response.NewResponse(200, output, nil)
	return ctx.Status(response.Code).JSON(response)

}
