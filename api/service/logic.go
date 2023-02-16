package service

import (
	"database/sql"
	"net/url"
	"regexp"
	"time"

	"github.com/fignocius/logic-api/api/model"
	"github.com/fignocius/logic-api/api/repository"
	"github.com/jmoiron/sqlx"
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

func (s Logic) UpInsert(logic model.Logic) (logicSave model.Logic, err error) {
	var (
		dbLogic model.Logic
		regex   = regexp.MustCompile(`\b[a-z]\b`)
	)

	if dbLogic, err = s.Repository.Get(logic.ID); err != nil {
		if err != sql.ErrNoRows {
			return
		}
		logic.ID = uuid.NewV4()
		logic.ExpresionCode = regex.ReplaceAllString(logic.Expression, "?")
		if err = s.Repository.Create(logic); err != nil {
			return
		}
		logicSave = logic
		return
	}
	dbLogic.ExpresionCode = regex.ReplaceAllString(logic.Expression, "?")
	dbLogic.UpdatedAt = time.Now()
	if err = s.Repository.Update(dbLogic); err != nil {
		return
	}
	logicSave = dbLogic
	return
}
func (s Logic) Apply(id uuid.UUID, queryParams url.Values) (result bool, err error) {
	var params []interface{}
	for _, v := range queryParams {
		params = append(params, v[0])
	}
	result, err = s.Repository.Apply(id, params)
	return
}

func (s Logic) Delete(id uuid.UUID) (err error) {
	err = s.Repository.Delete(id)
	return
}
