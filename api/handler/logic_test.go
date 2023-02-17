package handler

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fignocius/logic-api/api/model"
	"github.com/fignocius/logic-api/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogicHandler(t *testing.T) {
	var (
		mockedLogicService = &mocks.LogicService{}
		mockedLogicHandler = Logic{
			mockedLogicService,
		}
	)

	t.Run("testing list logic handler error", func(t *testing.T) {

		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/expressions", nil)
		recorder := httptest.NewRecorder()
		context := e.NewContext(request, recorder)

		mockedLogicService.Mock = mock.Mock{}
		mockedLogicService.On("List").Return([]model.Logic{}, sql.ErrNoRows)

		err := mockedLogicHandler.List(context)
		assert.Error(t, err)
		assert.EqualError(t, err, "code=500, message=Internal Server Error")
	})

	t.Run("testing upsert logic handler validation error", func(t *testing.T) {

		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/expressions", nil)
		recorder := httptest.NewRecorder()
		context := e.NewContext(request, recorder)

		err := mockedLogicHandler.Upsert(context)
		assert.Error(t, err)
		assert.EqualError(t, err, "code=400, message=expression: cannot be blank.")
	})
}
