package service

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/fignocius/logic-api/api/model"
	"github.com/fignocius/logic-api/mocks"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogicService(t *testing.T) {
	mockedLogicRepository := &mocks.LogicRepository{}
	logicMockedService := Logic{
		Repository: mockedLogicRepository,
	}
	t.Run("apply success test", func(t *testing.T) {
		var (
			id          = uuid.NewV4()
			queryParams = url.Values{"x": []string{"1"}, "y": []string{"0"}}
			params      = []interface{}{"1", "0"}
			received    = model.Logic{
				ID:             id,
				Expression:     "(x AND y)",
				ExpressionCode: "(? AND ?)",
			}
			err error
		)
		mockedLogicRepository.Mock = mock.Mock{}
		mockedLogicRepository.On("Get", id).Return(received, nil)
		mockedLogicRepository.On("Apply", received.ExpressionCode, params).Return(true, nil)
		result, err := logicMockedService.Apply(id, queryParams)
		assert.NoError(t, err)
		assert.Equal(t, true, result)
	})

	t.Run("apply database error test", func(t *testing.T) {
		var (
			id          = uuid.NewV4()
			queryParams = url.Values{"x": []string{"1"}, "y": []string{"0"}}
			params      = []interface{}{"1", "0"}
			received    = model.Logic{
				ID:             id,
				Expression:     "(x AND y)",
				ExpressionCode: "(? AND ?)",
			}
			err error
		)
		mockedLogicRepository.Mock = mock.Mock{}
		mockedLogicRepository.On("Get", id).Return(received, nil)
		mockedLogicRepository.On("Apply", received.ExpressionCode, params).Return(true, fmt.Errorf("database_error"))
		_, err = logicMockedService.Apply(id, queryParams)
		assert.Error(t, err)
		assert.EqualError(t, err, "database_error")
	})

	t.Run("apply parameters quantity error test", func(t *testing.T) {
		var (
			id          = uuid.NewV4()
			queryParams = url.Values{"x": []string{"1"}, "y": []string{"0"}, "z": []string{"0"}}
			received    = model.Logic{
				ID:             id,
				Expression:     "(x AND y)",
				ExpressionCode: "(? AND ?)",
			}
			err error
		)
		mockedLogicRepository.Mock = mock.Mock{}
		mockedLogicRepository.On("Get", id).Return(received, nil)
		_, err = logicMockedService.Apply(id, queryParams)
		assert.Error(t, err)
		assert.EqualError(t, err, "code=400, message=parameters error")
	})

	t.Run("apply parameters error test", func(t *testing.T) {
		var (
			id          = uuid.NewV4()
			queryParams = url.Values{"x": []string{"1"}, "z": []string{"0"}}
			received    = model.Logic{
				ID:             id,
				Expression:     "(x AND y)",
				ExpressionCode: "(? AND ?)",
			}
			err error
		)
		mockedLogicRepository.Mock = mock.Mock{}
		mockedLogicRepository.On("Get", id).Return(received, nil)
		_, err = logicMockedService.Apply(id, queryParams)
		assert.Error(t, err)
		assert.EqualError(t, err, "code=400, message=variable not found in this expression")
	})

	t.Run("upsert success test", func(t *testing.T) {
		var (
			id       = uuid.NewV4()
			received = model.Logic{
				ID:         id,
				Expression: "(x AND y)",
			}
			processed = model.Logic{
				ID:             id,
				Expression:     "(x AND y)",
				ExpressionCode: "(? AND ?)",
			}
			err error
		)
		mockedLogicRepository.Mock = mock.Mock{}
		mockedLogicRepository.On("Get", id).Return(received, nil)
		mockedLogicRepository.On("Update", processed).Return(nil)
		result, err := logicMockedService.Upsert(received)
		assert.NoError(t, err)
		assert.Equal(t, processed, result)
	})

	t.Run("upsert error test", func(t *testing.T) {
		var (
			id       = uuid.NewV4()
			received = model.Logic{
				ID:         id,
				Expression: "(x AND y)",
			}
			processed = model.Logic{
				ID:             id,
				Expression:     "(x AND y)",
				ExpressionCode: "(? AND ?)",
			}
			err error
		)
		mockedLogicRepository.Mock = mock.Mock{}
		mockedLogicRepository.On("Get", id).Return(received, nil)
		mockedLogicRepository.On("Update", processed).Return(fmt.Errorf("database_error"))
		_, err = logicMockedService.Upsert(received)
		assert.Error(t, err)
		assert.EqualError(t, err, "database_error")
	})
}
