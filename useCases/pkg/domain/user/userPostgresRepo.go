package user

import (
	"database/sql"
	"errors"

	uuid "github.com/satori/go.uuid"
)

type repo struct {
	store *sql.DB
	u     string
}

//NewPostgressRepo accepts db handler
func NewPostgressRepo(db *sql.DB) Repository {
	return &repo{
		store: db,
	}
}

func (r *repo) FindByID(ID string) (*User, error) {
	query := "SELECT first_name, last_name, email, id FROM app_user WHERE id = $1"
	rows, err := r.store.Query(query, ID)

	if err != nil {
		panic(err)
	}

	var u User
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&u.FirstName, &u.LastName, &u.Email, &u.ID)
	}

	if uuid.Equal(u.ID, uuid.NullUUID{}.UUID) {
		return nil, errors.New("Record not found")
	}

	return &u, nil
}

func (r *repo) Create(u *User) error {
	query := `INSERT INTO app_user(Id, first_name,last_name, email) VALUES($1 ,$2 ,$3, $4)`
	_, err := r.store.Exec(query, u.ID, u.FirstName, u.LastName, u.Email)

	return err
}

func (r *repo) DeleteUserByID(ID string) error {
	query := `DELETE FROM app_user WHERE id = $1`
	result, err := r.store.Exec(query, ID)

	rowSize, _ := result.RowsAffected()
	if rowSize == 0 {
		return errors.New("Record not found")
	}
	return err
}

func (r *repo) UpdateUserByID(uuid string) error {
	return nil
}
