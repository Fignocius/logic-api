package repository

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/fignocius/logic-api/api/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type Logic struct {
	db *sqlx.DB
}

func NewLogicsRepository(db *sqlx.DB) LogicRepository {
	return &Logic{
		db: db,
	}
}

func (r Logic) Apply(expressionCode string, queryParams []interface{}) (result bool, err error) {

	statement, _, err := psql.Select(expressionCode).ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	err = stmt.Get(&result, queryParams...)
	if err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (r Logic) List() (logic []model.Logic, err error) {
	query := psql.Select("*").
		From("logics")

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Select(&logic, args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (r Logic) Create(logic model.Logic) (err error) {
	query := psql.Insert("logics").
		Columns(
			"id",
			"expression",
			"expression_code",
		).
		Values(
			logic.ID,
			logic.Expression,
			logic.ExpressionCode,
		)
	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if _, err = stmt.Exec(args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (r Logic) Update(logic model.Logic) (err error) {
	query := psql.Update("logics").
		Set("expression", logic.Expression).
		Set("expression_code", logic.ExpressionCode).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": logic.ID})
	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if _, err = stmt.Exec(args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (r Logic) Get(id uuid.UUID) (logic model.Logic, err error) {

	query := psql.Select("*").
		From("logics").
		Where(sq.Eq{"id": id})

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Get(&logic, args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (r Logic) Delete(id uuid.UUID) (err error) {

	query := psql.Delete("logics").
		Where(sq.Eq{"id": id})

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if _, err = stmt.Exec(args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}
