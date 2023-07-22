package usecase

import (
	"time"

	"github.com/ivmello/finanzen-api/internal/adapter/database/mysql/repository"
	"github.com/ivmello/finanzen-api/internal/application/transaction/domain"
	"github.com/ivmello/finanzen-api/internal/application/transaction/dto"
)

type CreateTransactionUseCaseInterface interface {
	Execute(inputDto dto.CreateTransactionInputDto) (dto.CreateTransactionOutputDto, error)
}

type usecase struct {
	repository repository.TransactionMysqlRepositoryInterface
}

func NewCreateTransactionUseCase(repository repository.TransactionMysqlRepositoryInterface) CreateTransactionUseCaseInterface {
	return &usecase{
		repository,
	}
}

func (u *usecase) Execute(inputDto dto.CreateTransactionInputDto) (dto.CreateTransactionOutputDto, error) {
	transactionDate, err := time.Parse(time.RFC3339, inputDto.TransactionDate)
	if err != nil {
		return dto.CreateTransactionOutputDto{}, err
	}

	transactionDomain, err := domain.NewTransactionDomain(inputDto.Amount, inputDto.TransactionType, transactionDate, inputDto.Description)
	if err != nil {
		return dto.CreateTransactionOutputDto{}, err
	}

	_, err = u.repository.CreateTransaction(*transactionDomain)
	if err != nil {
		return dto.CreateTransactionOutputDto{}, err
	}

	return dto.CreateTransactionOutputDto{
		ID:              transactionDomain.ID,
		Amount:          transactionDomain.Amount,
		Description:     transactionDomain.Description,
		TransactionType: transactionDomain.TransactionType,
		TransactionDate: transactionDomain.TransactionDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}
