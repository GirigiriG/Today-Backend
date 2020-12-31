package project

import (
	"database/sql"
	"fmt"
)

const (
	allFields = `id, project_name, created_by, status, created_date, last_modified_date, sprint_id, description, percent_complete`
)

//NewProjectRepository holds repo db connection
type newProjectRepository struct {
	database *sql.DB
}

func NewProjectRepo(db *sql.DB) Repository {
	return &newProjectRepository{
		database: db,
	}
}

//CreateNewProjejct create new project record
func (repo *newProjectRepository) CreateNewProjejct(p *Project) error {
	query := fmt.Sprintf(`INSERT INTO project 
		(%s)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`, allFields)

	_, err := repo.database.Exec(query, p.ID, p.ProjectName, p.CreatedBy,
		p.Status, p.CreatedDate, p.LastModifiedDate, p.SprintID,
		p.Description, p.PercentComplete)

	if err != nil {
		return err
	}

	return nil
}

//GetProjectByID get a project by id
func (repo *newProjectRepository) GetProjectByID(ID string) (*Project, error) {
	query := fmt.Sprintf(`SELECT * FROM project WHERE id = $1`)

	rows, err := repo.database.Query(query, &ID)

	if err != nil {
		return nil, err
	}

	p := &Project{}

	for rows.Next() {
		rows.Scan(&p.ID, &p.ProjectName, &p.CreatedBy, &p.Status, &p.CreatedDate, &p.LastModifiedDate,
			&p.SprintID, &p.Description, &p.PercentComplete)
	}

	return p, nil
}

//UpdateProjectByID update a project record by id
func (repo *newProjectRepository) UpdateProjectByID(newProjectRecord *Project) (*Project, error) {
	return nil, nil
}

//DeleteProjectByID delete project record by id
func (repo *newProjectRepository) DeleteProjectByID(ID string) error {
	query := "`DELETE FROM project WHERE id = $1`"
	_, err := repo.database.Exec(query, ID)
	if err != nil {
		return err
	}
	return nil
}
