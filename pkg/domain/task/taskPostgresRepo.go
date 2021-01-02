package task

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

const (
	NoRecordFound = "No record found"
)

type repo struct {
	database *sql.DB
}

//NewTaskRepo instance of new repo
func NewTaskRepo(db *sql.DB) Repository {
	return &repo{
		database: db,
	}
}

func (repo *repo) Create(t *Task) error {
	query := `INSERT INTO task 
	(id, task_name, owner_id, created_date, last_modified_date, status, created_by, project_id, estimate, remaining, sprint_id)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := repo.database.Exec(query,
		t.ID, t.TaskName, t.OwnerID, t.CreatedDate,
		t.LastModifiedDate, t.Status, t.CreatedBy,
		t.ProjectID, t.Estimate, t.Remaining,
		t.SprintID)

	if err != nil {
		return err
	}
	return nil
}

//DeleteTaskByID delete record by id
func (repo *repo) DeleteByID(ID string) error {
	fmt.Println(ID)
	query := `DELETE FROM task WHERE id = $1`
	Results, err := repo.database.Exec(query, ID)
	n, _ := Results.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("Record not found.\n")
	}
	return nil
}

func (repo *repo) Update(t *Task) error {
	query := `
		UPDATE task
		SET task_name=$1, owner_id=$2, last_modified_date=$3, status=$4, project_id=$5, estimate=$6, remaining=$7, sprint_id=$8`
	results, err := repo.database.Exec(query,
		t.TaskName, t.OwnerID, t.LastModifiedDate,
		t.Status, t.ProjectID, t.Estimate, t.Remaining,
		t.SprintID)

	if err != nil {
		return err
	}

	n, _ := results.RowsAffected()
	if n == 0 {
		return fmt.Errorf("Record not found.\n")
	}
	return nil
}

func (repo *repo) FindByID(ID string) (*Task, error) {

	query := `SELECT id FROM task WHERE id = $1`

	rows, err := repo.database.Query(query, ID)
	if err != nil {
		return nil, errors.New(NoRecordFound)
	}

	t := &Task{}

	for rows.Next() {
		rows.Scan(&t.ID, &t.TaskName, &t.OwnerID, &t.CreatedDate, &t.LastModifiedDate,
			&t.Status, &t.CreatedBy, &t.ProjectID, &t.Estimate, &t.Remaining, &t.SprintID)
	}
	return t, nil

}

//FindAllTaskByProjectID find all test related to a project by projectid
func (repo *repo) FindAllByProjectID(IDs []string) ([]Task, error) {
	var tasks []Task
	query := `SELECT task_name FROM task WHERE project_id = ANY($1);`
	rows, err := repo.database.Query(query, pq.Array(IDs))
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.TaskName)
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		return nil, errors.New(NoRecordFound)
	}
	return tasks, nil
}
