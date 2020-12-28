package task

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"
)

//Task Object
type Task struct {
	ID               uuid.UUID
	TaskName         string
	OwnerID          string
	CreatedDate      time.Time
	LastModifiedDate time.Time
	Status           string
	CreatedBy        string
	ProjectID        string
	Estimate         int
	Remaining        int
	SprintID         string
}

//NewTask create a new task record type
func NewTask(t *Task) (*Task, error) {
	if err := validateNewTask(t); err != nil {
		return nil, err
	}

	ID := tools.CreateUUID()
	return &Task{
		ID:               ID,
		TaskName:         t.TaskName,
		OwnerID:          t.OwnerID,
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
		Status:           t.Status,
		CreatedBy:        t.CreatedBy,
		ProjectID:        t.ProjectID,
		Estimate:         t.Estimate,
		Remaining:        t.Remaining,
		SprintID:         t.SprintID,
	}, nil
}

func validateNewTask(t *Task) error {
	if t.TaskName == ""  {
		return errors.New("Please provide a task name")
	}
	if t.OwnerID == "" {
		return errors.New("Pleasae provided owner id")
	}
	if t.Status == "" {
		return errors.New("Pleasae provided status")
	}
	if t.ProjectID == "" {
		return errors.New("Pleasae provided project id")
	}
	if t.SprintID == "" {
		return errors.New("Pleasae provided sprint id")
	}
	return nil
}
