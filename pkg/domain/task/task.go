package task

import (
	"errors"
	"time"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"
)

const (
	NameRequired                    = "Please provide a task name"
	OwnerIDRequired                 = "Pleasae provided owner id"
	StatusRequired                  = "Pleasae provided status"
	ProjectIDRequired               = "Pleasae provided project id"
	SprintIDRequired                = "Pleasae provided sprint id"
	RemainingHoursMustBeEqaulToZero = "This task still has remaining hours"
)

//Task Object
type Task struct {
	ID               string
	TaskName         string
	OwnerID          string
	CreatedDate      time.Time
	LastModifiedDate time.Time
	Status           string
	CreatedBy        string
	ProjectID        string
	Estimate         int32
	Remaining        int32
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

	if t.TaskName == "" {
		return errors.New(NameRequired)
	}

	if t.OwnerID == "" {
		return errors.New(OwnerIDRequired)
	}

	if t.Status == "" {
		return errors.New(StatusRequired)
	}

	if t.ProjectID == "" {
		return errors.New(ProjectIDRequired)
	}
	if t.SprintID == "" {
		return errors.New(SprintIDRequired)
	}

	if t.Status == "Closed" && t.Remaining > 0 {
		return errors.New(RemainingHoursMustBeEqaulToZero)
	}
	return nil
}
