package user

import (
	"database/sql"
	"errors"
)

type repo struct {
	database *sql.DB
}

//NewPostgressRepo accepts db handler
func NewPostgressRepo(db *sql.DB) Repository {
	return &repo{
		database: db,
	}
}

func (r *repo) FindByID(ID string) (*User, error) {
	query := `SELECT id,first_name,last_name,email FROM app_user WHERE id =$1`
	rows, err := r.database.Query(query, ID)
	if err != nil {
		return nil, err
	}

	var u User

	for rows.Next() {
		rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	}
	return &u, nil
}

func (r *repo) Create(u *User) error {
	query := `INSERT INTO app_user(Id, first_name,last_name, email) VALUES($1 ,$2 ,$3, $4)`
	results, err := r.database.Exec(query, u.ID, u.FirstName, u.LastName, u.Email)
	n, _ := results.RowsAffected()
	if n == 0 {
		return errors.New("Record not found")
	}
	return err
}

func (r *repo) DeleteByID(uuid string) error {
	query := `DELETE FROM app_user WHERE id = $1`
	result, err := r.database.Exec(query, uuid)

	rowSize, _ := result.RowsAffected()
	if rowSize == 0 {
		return errors.New("Record not found")
	}
	return err
}

func (r *repo) UpdateByID(u *User) error {
	query := `UPDATE app_user
	SET first_name= $1, last_name= $2,email= $3
	WHERE id = $4
	`
	results, err := r.database.Exec(query, u.FirstName, u.LastName, u.Email, u.ID)
	if err != nil {
		return err
	}

	n, _ := results.RowsAffected()
	if n == 0 {
		return errors.New("Record not found")
	}
	return nil

}
