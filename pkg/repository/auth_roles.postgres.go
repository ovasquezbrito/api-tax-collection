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

func (r *RolePostgres) CreateRole(ctx context.Context, role entity.Role) (int, error) {
	var id int
	query := fmt.Sprintf(
		`
			INSERT INTO %s (
				role_name, role_nivel 
			) 
			values ($1, $2) 
			RETURNING id`,
		roleTable,
	)

	_ = r.db.QueryRowContext(ctx, query, role.RoleName, role.RoleNivel).Scan(&id)
	return id, nil
}

func (r *RolePostgres) GetAll(ctx context.Context, queryp entity.QueryParameter) ([]entity.Role, int, error) {
	var lists []entity.Role
	var total int

	offset := (queryp.Page - 1) * queryp.Limit
	search_field_one := "'%" + queryp.Search + "%'"

	query := fmt.Sprintf(`SELECT id, role_name, role_nivel, status 
						  FROM %s
						  WHERE role_name like %s
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

func (r *RolePostgres) GetRoleByName(ctx context.Context, rolName string) (*entity.Role, error) {
	item := &entity.Role{}
	query := fmt.Sprintf(`SELECT id, TRIM(role_name), role_nivel, status 
		FROM %s
		WHERE TRIM(role_name) = $1`,
		roleTable,
	)

	err := r.db.GetContext(ctx, item, query, rolName)
	if err != nil {
		return item, errors.New("error al realizar la consulta")
	}
	fmt.Println("esto viene de la tabla", item.Id)
	return item, nil
}

func (r *RolePostgres) GetById(ctx context.Context, idRol int) (*entity.Role, error) {
	item := &entity.Role{}
	query := fmt.Sprintf(`SELECT role_name, role_nivel, status 
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

func (r *RolePostgres) DeleteById(ctx context.Context, idRol int) (int64, error) {
	var id int64
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, roleTable)

	a, err := r.db.ExecContext(ctx, query, idRol)
	if err != nil {
		return 0, errors.New("error al eliminar el registro")
	}

	id, err = a.RowsAffected()
	if err != nil {
		return 0, err
	}

	return id, nil
}
