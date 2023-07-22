package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type InputDto struct {
	Amount          int64  `json:"amount" validate:"required,min=1"`
	TransactionType string `json:"transaction_type" validate:"required"`
}

func TestValidateStruct(t *testing.T) {
	t.Run("should return a nil error", func(t *testing.T) {
		inputDto := InputDto{
			Amount:          100,
			TransactionType: "CREDIT",
		}
		err := ValidateStruct(inputDto)

		assert.Nil(t, err)
	})

	t.Run("should return a nil error", func(t *testing.T) {
		inputDto := InputDto{
			Amount:          -1,
			TransactionType: "CREDIT",
		}
		err := ValidateStruct(inputDto)

		assert.Equal(t, "Required min size is 1", err.Causes[0].Message)
	})
}
