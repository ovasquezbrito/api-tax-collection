package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ovasquezbrito/tax-collection/pkg/entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(ctx context.Context, user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		`
			INSERT INTO %s (
				first_last_name, 
				email, 
				password,
				avatar_user
			) 
			values ($1, $2, $3, $4) 
			RETURNING id`,
		usersTable,
	)

	row := r.db.QueryRow(
		query,
		user.FirstLast,
		user.Email,
		user.Password,
		user.AvatarUser,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(ctx context.Context, email, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf(
		`
		SELECT us.id, us.first_last_name as name, us.email, us.password, us.avatar_user 
        FROM %s AS us
     	WHERE trim(us.email)=$1 AND trim(us.password)=$2
	`,
		usersTable,
	)
	err := r.db.Get(&user, query, email, password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthPostgres) UpdateUser(ctx context.Context, idUser int, user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		`
			UPDATE %s SET
				first_last_name = $1,
				email = $2,
				password = $3,
				avatar_user = $4
			WHERE id = $5 
			RETURNING id
		`,
		usersTable,
	)

	row := r.db.QueryRow(
		query,
		user.FirstLast,
		user.Email,
		user.Password,
		user.AvatarUser,
		idUser,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserById(ctx context.Context, idUser int) (*entity.User, error) {
	u := &entity.User{}
	query := fmt.Sprintf(
		`
			SELECT us.id, us.first_last_name as name, us.email, us.username, us.password, us.avatar_user
     	FROM %s AS us
     	WHERE us.id = $1
		`,
		usersTable,
	)

	err := r.db.Get(&u, query, idUser)
	if err != nil {
		return u, err //errors.New("error al realizar la consulta")
	}
	return u, nil
}

func (r *AuthPostgres) GetUserByUserName(ctx context.Context, email string) (int, error) {
	var user entity.User
	query := fmt.Sprintf(
		`
			SELECT us.first_last_name as name, us.email, us.password, us.avatar_user
     	FROM %s AS us
     	WHERE trim(us.email) = $1
		`,
		usersTable,
	)

	err := r.db.Get(&user, query, email)
	if err != nil {
		return 0, errors.New("no hubo resultado")
	}
	return 1, errors.New("ya existe un usuario con este nombre")
}

func (r *AuthPostgres) GetUserByUserEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	query := fmt.Sprintf(
		`
			SELECT us.id, us.first_last_name as name, us.email, us.password, us.avatar_user
     	FROM %s AS us
     	WHERE trim(us.email) = $1
		`,
		usersTable,
	)

	err := r.db.Get(u, query, email)
	if err != nil {
		return u, errors.New("no hubo resultado")
	}
	return u, nil
}

func (r *AuthPostgres) GetMenuOptionAll(ctx context.Context, idUser int) ([]entity.RoleUser, error) {
	var lists []entity.RoleUser

	query := fmt.Sprintf(
		`
		SELECT menu.name_opcion, menu.icon, menu.componente_uri, menu.page_url, menu.orderby, menu.type_opcion, 
	           rou.nivel_opcion, menu."id"
        FROM %s AS rou
	         INNER JOIN %s AS menu ON  rou.fk_menu_opcion_detalle = menu."id"
        WHERE rou.fk_id_user = $1
        ORDER BY menu.orderby ASC
		 `,
		roleUserTable,
		menuOpcionDetalleTable,
	)

	err := r.db.Select(&lists, query, idUser)
	if err != nil {
		return nil, err
	}

	return lists, nil

}
