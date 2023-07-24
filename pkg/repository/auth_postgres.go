package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	baseapp "github.com/ovasquezbrito/tax-collection"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user baseapp.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		`
			INSERT INTO %s (
				first_last, 
				email, 
				password,
				uri_img
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
		user.UriImg,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (baseapp.User, error) {
	var user baseapp.User
	query := fmt.Sprintf(
		`
		SELECT us.id, us.first_last as name, us.email, us.password, us.uri_img 
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

func (r *AuthPostgres) UpdateUser(idUser int, user baseapp.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		`
			UPDATE %s SET
				first_last = $1,
				email = $2,
				password = $3,
				uri_img = $4
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
		user.UriImg,
		idUser,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUserById(idUser int) (baseapp.User, error) {
	var user baseapp.User
	query := fmt.Sprintf(
		`
			SELECT us.id, us.first_last as name, us.email, us.username, us.password, us.uri_img
     	FROM %s AS us
     	WHERE us.id = $1
		`,
		usersTable,
	)

	err := r.db.Get(&user, query, idUser)
	if err != nil {
		return user, err //errors.New("error al realizar la consulta")
	}
	return user, nil
}

func (r *AuthPostgres) GetUserByUserName(email string) (int, error) {
	var user baseapp.User
	query := fmt.Sprintf(
		`
			SELECT us.first_last as name, us.email, us.password, us.uri_img
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

func (r *AuthPostgres) GetUserByUserEmail(email string) (baseapp.User, error) {
	var user baseapp.User
	query := fmt.Sprintf(
		`
			SELECT us.id, us.first_last as name, us.email, us.password, us.uri_img
     	FROM %s AS us
     	WHERE trim(us.email) = $1
		`,
		usersTable,
	)

	err := r.db.Get(&user, query, email)
	if err != nil {
		return user, errors.New("no hubo resultado")
	}
	return user, nil
}

func (r *AuthPostgres) GetMenuOptionAll(idUser int) ([]baseapp.RoleUser, error) {
	var lists []baseapp.RoleUser

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
