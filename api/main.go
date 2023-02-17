package main

import (
	"fmt"
	"os"

	"github.com/fignocius/logic-api/api/handler"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	defer db.Close()

	logicHandler := handler.NewLogicHandler(db)

	e := echo.New()
	e.Use(mw.KeyAuth(func(key string, c echo.Context) (bool, error) {
		fmt.Println(key)
		return key == os.Getenv("KEY_AUTH_API"), nil
	}))
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/evaluate/:expression_id", logicHandler.Apply)
	e.DELETE("/expressions/:expression_id", logicHandler.Delete)
	e.GET("/expressions", logicHandler.List)
	e.POST("/expressions", logicHandler.Upsert)
	e.Logger.Fatal(e.Start(":8080"))
}
