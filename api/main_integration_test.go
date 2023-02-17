package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/fignocius/logic-api/api/handler"
	"github.com/fignocius/logic-api/api/model"
	"github.com/gavv/httpexpect"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

var dbConn *sqlx.DB

func TestMain(m *testing.M) {
	var err error
	dbConn, err = sqlx.Connect("postgres", "postgres://postgres:postgres@localhost:5433/logic-api-test?sslmode=disable")
	if err != nil {
		log.Panicf("Could not connect to database: %s", err.Error())
	}
	driver, err := postgres.WithInstance(dbConn.DB, &postgres.Config{})
	if err != nil {
		log.Panicf("driver err: %v", err)
	}
	dbMigrate, err := migrate.NewWithDatabaseInstance("file://../migrations", "logic-api-test", driver)
	if err != nil {
		log.Panicf("migrate err: %v", err)
	}
	_ = dbMigrate.Up()
	initData()

	code := m.Run()
	_ = dbConn.Close()
	os.Exit(code)
}

func initData() {
	file, _ := os.Open("./testdata/db_inserts.sql")
	defer file.Close()
	bs, _ := io.ReadAll(file)
	inserts := string(bs)

	_, err := dbConn.Exec(inserts)
	if err != nil {
		log.Panicf("Could not insert data: %s", err.Error())
	}
}

func TestGetOlistShippings(t *testing.T) {
	handler := handler.NewLogicHandler(dbConn)

	e := echo.New()
	e.GET("/evaluate/:expression_id", handler.Apply)
	e.DELETE("/expressions/:expression_id", handler.Delete)
	e.GET("/expressions", handler.List)
	e.POST("/expressions", handler.Upsert)
	defer e.Close()

	server := httptest.NewServer(e.Server.Handler)
	defer server.Close()
	ex := httpexpect.New(t, server.URL)

	t.Run("List Success", func(t *testing.T) {
		obj := ex.GET("/expressions").
			Expect().
			Status(http.StatusOK).
			JSON().
			Array()

		obj.Element(0).Object().Value("expression").Equal("x OR y")
		obj.Element(1).Object().Value("expression").Equal("x AND y")
		obj.Element(2).Object().Value("expression").Equal("(x AND y) OR z")
	})

	t.Run("evaluate true Success", func(t *testing.T) {
		obj := ex.GET("/evaluate/8f1196dc-a2f1-4667-81f2-99023cf7c5ea").
			WithQuery("x", "1").
			WithQuery("y", "1").
			Expect().
			Status(http.StatusOK)

		obj.Body().Equal("true\n")
	})

	t.Run("evaluate false Success", func(t *testing.T) {
		obj := ex.GET("/evaluate/8f1196dc-a2f1-4667-81f2-99023cf7c5ea").
			WithQuery("x", "0").
			WithQuery("y", "1").
			Expect().
			Status(http.StatusOK)

		obj.Body().Equal("false\n")
	})

	t.Run("update Success", func(t *testing.T) {
		obj := ex.POST("/expressions").
			WithJSON(model.Logic{
				ID:         uuid.FromStringOrNil("8f1196dc-a2f1-4667-81f2-99023cf7c5ea"),
				Expression: "z OR y",
			}).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object()

		obj.Value("expression").Equal("z OR y")
	})

	t.Run("delete Success", func(t *testing.T) {
		ex.DELETE("/expressions/f4c01f76-b202-4d13-826b-46cd6fd97249").
			Expect().
			Status(http.StatusNoContent)
	})
}
