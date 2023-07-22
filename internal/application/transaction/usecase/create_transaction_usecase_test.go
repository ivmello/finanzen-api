package usecase

import (
	"testing"
)

func TestCreateTransactionUseCase(t *testing.T) {
	t.Run("should create a new transaction", func(t *testing.T) {
		// db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		// if err != nil {
		// 	log.Fatal("error init mock ", err)
		// }
		// defer db.Close()

		// repo := repository.NewTransactionMysqlRepositoryMock(db)
		// service := NewCreateTransactionUseCase(repo)

		// transaction, _ := domain.NewTransactionDomain(100, "CREDIT", time.Now(), "Test")

		// mock.ExpectExec("INSERT INTO transactions (id, amount, transaction_type, transaction_date, description, created_at, updated_at)").WithArgs(
		// 	transaction.ID,
		// 	transaction.Amount,
		// 	transaction.TransactionType,
		// 	transaction.TransactionDate,
		// 	transaction.Description,
		// 	transaction.CreatedAt,
		// 	transaction.UpdatedAt,
		// ).WillReturnResult(sqlmock.NewResult(1, 1))

		// inputDto := dto.CreateTransactionInputDto{
		// 	Amount:          transaction.Amount,
		// 	TransactionType: transaction.TransactionType,
		// 	TransactionDate: "2023-07-21T15:34:05Z",
		// 	Description:     transaction.Description,
		// }

		// output, err := service.Execute(inputDto)

		// assert.Nil(t, err)
		// assert.Equal(t, transaction.ID, output.ID)

		t.Skip()
	})
}
