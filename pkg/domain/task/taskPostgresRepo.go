package task

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type repo struct {
	store *sql.DB
}

func NewTaskRepo(db *sql.DB) Repository {
	return &repo{
		store: db,
	}
}

func (repo *repo) CreateTask(t *Task) error {
	query := `INSERT INTO task 
	(id, task_name, owner_id, created_date, last_modified_date, status, created_by, project_id, estimate, remaining, sprint_id)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := repo.store.Exec(query,
		t.ID, t.TaskName, t.OwnerID, t.CreatedDate,
		t.LastModifiedDate, t.Status, t.CreatedBy,
		t.ProjectID, t.Estimate, t.Remaining,
		t.SprintID)

	if err != nil {
		return err
	}
	return nil
}

func (rep *repo) DeleteTaskByID(ID string) error {
	return nil
}
func (repo *repo) FindTaskByID(ID string) (*Task, error) {
	return nil, nil
}

func (repo *repo) FindAllTaskByProjectID(IDs []string) (*[]Task, error) {
	var tasks []Task
	query := `SELECT task_name FROM task WHERE project_id = ANY($1);`
	rows, err := repo.store.Query(query, pq.Array(IDs))
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.TaskName)
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		return nil, errors.New("No record found")
	}
	return &tasks, nil
}
