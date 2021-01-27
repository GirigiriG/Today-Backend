package project

import (
	"database/sql"
	"fmt"
)

const (
	allFields = `
	id,
	name,
	created_by,
	status,
	created_date,
	last_modified_date,
	sprint_id,
	description,
	percent_complete`
)

//NewProjectRepository holds repo db connection
type newProjectRepository struct {
	database *sql.DB
}

//NewProjectRepo : requires database driver
func NewProjectRepo(db *sql.DB) Repository {
	return &newProjectRepository{
		database: db,
	}
}

//CreateNewProjejct create new project record
func (repo *newProjectRepository) Create(p *Project) error {
	query := fmt.Sprintf(`INSERT INTO project 
		(%s)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`, allFields)

	_, err := repo.database.Exec(query,
		p.ID,
		p.ProjectName,
		p.CreatedBy,
		p.Status,
		p.CreatedDate,
		p.LastModifiedDate,
		p.SprintID,
		p.Description,
		p.PercentComplete,
	)

	if err != nil {
		return err
	}

	return nil
}

//GetByID get a project by id
func (repo *newProjectRepository) FindByID(ID string) (*Project, error) {
	query := fmt.Sprintf(`SELECT * FROM project WHERE id = $1`)

	rows, err := repo.database.Query(query, &ID)

	if err != nil {
		return nil, err
	}

	p := &Project{}

	for rows.Next() {
		rows.Scan(
			&p.ID,
			&p.ProjectName,
			&p.CreatedBy,
			&p.Status,
			&p.CreatedDate,
			&p.LastModifiedDate,
			&p.SprintID,
			&p.Description,
			&p.PercentComplete,
		)
	}

	return p, nil
}

//UpdateByID update a project record by id
func (repo *newProjectRepository) UpdateByID(p *Project) error {
	query := `
		UPDATE project 
		SET 
			name=$1,
			status=$2,
			last_modified_date=$3,
			description=$4,
			percent_complete=$5,
			sprint_id=$6
		WHERE id=$7`

	results, err := repo.database.Exec(query,
		p.ProjectName,
		p.Status,
		p.LastModifiedDate,
		p.Description,
		p.PercentComplete,
		p.SprintID,
		p.ID,
	)

	if err != nil {
		return err
	}

	n, _ := results.RowsAffected()
	if n == 0 {
		return fmt.Errorf("Record not found")
	}
	return nil
}

//DeleteByID delete project record by id
func (repo *newProjectRepository) DeleteByID(ID string) error {
	query := `DELETE FROM project WHERE id = $1`
	results, err := repo.database.Exec(query, ID)
	if err != nil {
		return err
	}
	n, _ := results.RowsAffected()
	if n == 0 {
		return fmt.Errorf("record not found")
	}

	return nil
}
