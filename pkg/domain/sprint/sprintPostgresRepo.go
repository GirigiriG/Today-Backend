package sprint

import (
	"database/sql"
	"errors"
)

//Repo holds database
type Repo struct {
	database *sql.DB
}

//NewSprintRepositroy : create new sprint repo give a db
func NewSprintRepositroy(db *sql.DB) *Repo {
	return &Repo{
		database: db,
	}
}

//Create : creates new sprint record
func (repo *Repo) Create(s *Sprint) error {
	query := `INSERT INTO sprint (id, name, start_date, end_date)
		VALUES($1,$2,$3,$4)`
	results, err := repo.database.Exec(query, s.ID, s.SprintName, s.StartDate, s.EndDate)
	if err != nil {
		return err
	}

	n, _ := results.RowsAffected()
	if n == 0 {
		return errors.New("Record not found")
	}

	return nil
}

//Update : Update sprint record
func (repo *Repo) Update(s *Sprint) error {
	query := `UPDATE sprint
		SET name= $1, start_date= $2, end_date= $3
		WHERE id = $4`
	results, err := repo.database.Exec(query, s.SprintName, s.StartDate, s.EndDate, s.ID)
	if err != nil {
		return err
	}

	n, _ := results.RowsAffected()
	if n == 0 {
		return errors.New("Record not found")
	}
	return nil
}

//DeleteByID : find single sprint record by id
func (repo *Repo) DeleteByID(ID string) error {
	query := `DELETE FROM sprint WHERE id = $1`
	results, err := repo.database.Exec(query, ID)
	if err != nil {
		return err
	}
	n, _ := results.RowsAffected()
	if n == 0 {
		return errors.New("Record not found")
	}
	return nil
}

//FindByID : find single sprint record by id
func (repo *Repo) FindByID(ID string) (*Sprint, error) {
	query := `SELECT id, name, start_date, end_date  FROM sprint WHERE id = $1`

	rows, err := repo.database.Query(query, ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	s := &Sprint{}
	for rows.Next() {
		rows.Scan(s.ID, s.SprintName, s.StartDate, s.EndDate)
	}
	return s, nil
}
