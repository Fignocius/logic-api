package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fignocius/logic-api/api/model"
	uuid "github.com/satori/go.uuid"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type LogicRepository interface {
	Apply(string, []interface{}) (bool, error)
	List() ([]model.Logic, error)
	Create(model.Logic) error
	Update(model.Logic) error
	Get(uuid.UUID) (model.Logic, error)
	Delete(uuid.UUID) error
}
