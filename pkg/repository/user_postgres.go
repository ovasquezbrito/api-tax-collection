package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ovasquezbrito/tax-collection/pkg/entity"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) UpdateRoleToUser(ctx context.Context, idRole, idUser int) error {

	query := fmt.Sprintf(
		`
			UPDATE %s SET fk_role = $1 WHERE id = $2 
		`,
		usersTable,
	)

	_, err := r.db.ExecContext(ctx, query, idRole, idUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) GetAll(ctx context.Context, filter entity.QueryParameter) ([]entity.UserResponse, int, error) {
	var lists []entity.UserResponse
	var total int

	offset := (filter.Page - 1) * filter.Limit
	search_field_one := "'%" + filter.Search + "%'"

	sqlQuery := fmt.Sprintf(
		`
			SELECT u.id, u.first_last_name,	u.email, u.avatar_user, u.status, u.isadmin, u.fk_role,
			u.created_at, u.updated_at
				FROM %s as u
				WHERE u.first_last_name like %s
				ORDER BY u.first_last_name LIMIT $1 OFFSET $2
			`,
		usersTable,
		search_field_one,
	)

	err := r.db.SelectContext(ctx, &lists, sqlQuery, filter.Limit, offset)
	if err != nil {
		return nil, 0, err
	}

	sqlQueryTotal := fmt.Sprintf("SELECT count(*) as total FROM %s", usersTable)
	err = r.db.GetContext(ctx, &total, sqlQueryTotal)
	if err != nil {
		return nil, 0, errors.New("error al realizar la consulta")
	}

	return lists, total, nil
}
