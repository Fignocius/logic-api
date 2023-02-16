package repository

import (
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

func (t Logic) Apply(id uuid.UUID, queryParams []interface{}) (result bool, err error) {
	var logic model.Logic

	query := psql.Select("*").
		From("logics").
		Where(sq.Eq{"id": id})

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := t.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Get(&logic, args...); err != nil {
		return
	}
	statement, _, err = psql.Select(logic.ExpresionCode).ToSql()
	if err != nil {
		return
	}
	stmt, err = t.db.Preparex(statement)
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
			logic.ExpresionCode,
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
		Set("expression_code", logic.ExpresionCode).
		Set("updated_at", logic.UpdatedAt).
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
