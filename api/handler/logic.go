package handler

import (
	"net/http"

	"github.com/fignocius/logic-api/api/model"
	"github.com/fignocius/logic-api/api/service"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type Logic struct {
	service.LogicService
}

func NewLogicHandler(db *sqlx.DB) LogicHandler {
	return &Logic{
		service.New(db),
	}
}

func (t Logic) List(c echo.Context) error {
	logics, err := t.LogicService.List()
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, logics)
}

func (t Logic) UpInsert(c echo.Context) error {
	request := model.Logic{}
	err := c.Bind(&request)
	if err != nil {
		return echo.ErrBadRequest
	}
	products, err := t.LogicService.UpInsert(request)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, products)
}

func (t Logic) Apply(c echo.Context) error {
	id, err := uuid.FromString(c.Param("expression_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	query := c.QueryParams()
	products, err := t.LogicService.Apply(id, query)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, products)
}

func (t Logic) Delete(c echo.Context) error {
	id, err := uuid.FromString(c.Param("expression_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	err = t.LogicService.Delete(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusNoContent, nil)
}
