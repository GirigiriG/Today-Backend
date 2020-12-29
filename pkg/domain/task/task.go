package task

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"
)

const (
	NAME_RQUIRED = "Please provide a task name"	
	OWNER_ID_REQUIRED = "Pleasae provided owner id"	
	STATUS_REQUIRED = "Pleasae provided status"	
	PROJECT_ID_REQUIRED = "Pleasae provided project id"	
	SPRINT_ID_REQUIRED = "Pleasae provided sprint id"	
	HAS_REMAINING_HOURS = "This task still has remaining hours"	
	
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
		return errors.New(NAME_RQUIRED)
	}
	if t.OwnerID == "" {
		return errors.New(OWNER_ID_REQUIRED)
	}
	if t.Status == "" {
		return errors.New(STATUS_REQUIRED)
	}
	if t.ProjectID == "" {
		return errors.New(PROJECT_ID_REQUIRED)
	}
	if t.SprintID == "" {
		return errors.New(SPRINT_ID_REQUIRED)
	}
	if t.Status == "Closed" && t.Remaining > 0 {
		return errors.New(HAS_REMAINING_HOURS)

	}
	return nil
}
