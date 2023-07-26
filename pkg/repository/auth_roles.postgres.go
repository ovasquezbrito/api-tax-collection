package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ovasquezbrito/tax-collection/pkg/entity"
)

type RolePostgres struct {
	db *sqlx.DB
}

func NewRolePostgres(db *sqlx.DB) *RolePostgres {
	return &RolePostgres{db: db}
}

func (r *RolePostgres) GetAll(ctx context.Context, queryp entity.QueryParameter) ([]entity.Role, int, error) {
	var lists []entity.Role
	var total int

	offset := (queryp.Page - 1) * queryp.Limit
	search_field_one := "'%" + queryp.Search + "%'"

	query := fmt.Sprintf(`SELECT id, name_role, nivel_role 
						  FROM %s
						  WHERE name_role like %s
						  ORDER BY name_role LIMIT $1 OFFSET $2`,
		roleTable,
		search_field_one,
	)

	err := r.db.Select(&lists, query, queryp.Limit, offset)
	if err != nil {
		return nil, 0, err
	}

	query = fmt.Sprintf("SELECT count(*) as total FROM %s", roleTable)
	err = r.db.Get(&total, query)
	if err != nil {
		return nil, 0, errors.New("error al realizar la consulta")
	}

	return lists, total, nil
}

func (r *RolePostgres) GetById(ctx context.Context, idRol int) (*entity.Role, error) {
	item := &entity.Role{}
	query := fmt.Sprintf(`SELECT name_role, nivel_role 
		FROM %s
		WHERE id = $1`,
		roleTable,
	)

	err := r.db.Get(item, query, idRol)
	if err != nil {
		return item, errors.New("error al realizar la consulta")
	}
	return item, nil
}
