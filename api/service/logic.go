package service

import (
	"database/sql"
	"net/http"
	"net/url"
	"regexp"

	"github.com/fignocius/logic-api/api/model"
	"github.com/fignocius/logic-api/api/repository"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type Logic struct {
	Repository repository.LogicRepository
}

func New(db *sqlx.DB) LogicService {
	return &Logic{
		Repository: repository.NewLogicsRepository(db),
	}
}

func (s Logic) List() (logics []model.Logic, err error) {

	logics, err = s.Repository.List()

	return
}

func (s Logic) Upsert(logic model.Logic) (logicSave model.Logic, err error) {
	var (
		dbLogic model.Logic
		regex   = regexp.MustCompile(`\b[a-z]\b`)
	)

	if dbLogic, err = s.Repository.Get(logic.ID); err != nil {
		if err != sql.ErrNoRows {
			return
		}
		logic.ID = uuid.NewV4()
		logic.ExpressionCode = regex.ReplaceAllString(logic.Expression, "?")
		if err = s.Repository.Create(logic); err != nil {
			return
		}
		logicSave = logic
		return
	}
	dbLogic.ExpressionCode = regex.ReplaceAllString(logic.Expression, "?")
	if err = s.Repository.Update(dbLogic); err != nil {
		return
	}
	logicSave = dbLogic
	return
}
func (s Logic) Apply(id uuid.UUID, queryParams url.Values) (result bool, err error) {
	var (
		logic     model.Logic
		variables []string
		params    []interface{}
		regex     = regexp.MustCompile(`\b[a-z]\b`)
	)

	if logic, err = s.Repository.Get(id); err != nil {
		return
	}

	variables = regex.FindAllString(logic.Expression, -1)

	if len(queryParams) != len(variables) {
		err = echo.NewHTTPError(http.StatusBadRequest, "parameters error")
		return
	}

	for _, v := range variables {
		if _, ok := queryParams[v]; ok != true {
			err = echo.NewHTTPError(http.StatusBadRequest, "variable not found in this expression")
			return
		}
		params = append(params, queryParams[v][0])
	}

	result, err = s.Repository.Apply(logic.ExpressionCode, params)
	return
}

func (s Logic) Delete(id uuid.UUID) (err error) {
	err = s.Repository.Delete(id)
	return
}
