package handler

import (
	"github.com/labstack/echo/v4"
)

// LogicHandler interface
type LogicHandler interface {
	List(c echo.Context) error
	Upsert(c echo.Context) error
	Apply(c echo.Context) error
	Delete(c echo.Context) error
}
