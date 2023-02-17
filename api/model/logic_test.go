package model

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestLogicModel(t *testing.T) {
	t.Run("validation success test", func(t *testing.T) {
		var (
			id       = uuid.NewV4()
			received = Logic{
				ID:         id,
				Expression: "(x AND y)",
			}
			err error
		)
		err = received.Validate()
		assert.NoError(t, err)
	})

	t.Run("validation error test", func(t *testing.T) {
		var (
			id       = uuid.NewV4()
			received = Logic{
				ID: id,
			}
			err error
		)
		err = received.Validate()
		assert.Error(t, err)
		assert.EqualError(t, err, "expression: cannot be blank.")
	})
}
