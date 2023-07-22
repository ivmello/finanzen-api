package exception

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewException(t *testing.T) {
	t.Run("should return a InternalServerError exception", func(t *testing.T) {
		err := NewInternalServerError("internal server error")

		assert.Equal(t, "internal server error", err.Message)
		assert.Equal(t, http.StatusInternalServerError, err.Code)
	})

	t.Run("should return a BadGatewayError exception", func(t *testing.T) {
		err := NewBadGatewayError("bad gateway error")

		assert.Equal(t, "bad gateway error", err.Message)
		assert.Equal(t, http.StatusBadGateway, err.Code)
	})

	t.Run("should return a BadRequestError exception", func(t *testing.T) {
		causes := []*ErrorCauses{
			{
				Field:   "amount",
				Message: "This field is required",
			},
		}
		err := NewBadRequestError("bad gateway error", causes)

		assert.Equal(t, "bad gateway error", err.Message)
		assert.Equal(t, "amount", err.Causes[0].Field)
		assert.Equal(t, "This field is required", err.Causes[0].Message)
		assert.Equal(t, http.StatusBadRequest, err.Code)
	})
}
