package response

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResponse(t *testing.T) {
	t.Run("should return a new response", func(t *testing.T) {
		data := map[string]string{
			"message": "success",
		}
		response := NewResponse(200, data, nil)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, true, response.Sucess)
		assert.Equal(t, data, response.Data)
		assert.Nil(t, response.Errors)
	})
}
