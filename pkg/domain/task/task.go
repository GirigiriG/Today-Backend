package task

import (
	"errors"
	"time"
)

const (
	//NameRequired : Name required error messsage
	NameRequired = "Please provide a task name"
	//OwnerIDRequired : Owuner id required error messsage
	OwnerIDRequired = "Pleasae provided owner id"
	//StatusRequired : Status required error messsage
	StatusRequired = "Pleasae provided status"
	//ProjectIDRequired : Project id required error messsage
	ProjectIDRequired = "Pleasae provided project id"
	//SprintIDRequired : Sprint id required error messsage
	SprintIDRequired = "Pleasae provided sprint id"
	//RemainingHoursMustBeEqaulToZero : Task must have 0 hours remaining error message
	RemainingHoursMustBeEqaulToZero = "This task still has remaining hours"
)

//Task Object
type Task struct {
	ID               string
	TaskName         string
	OwnerID          string
	OwnerName        string
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

	newRecord := &Task{
		ID:               t.ID,
		TaskName:         t.TaskName,
		OwnerID:          t.OwnerID,
		OwnerName:        t.OwnerName,
		CreatedDate:      time.Now(),
		LastModifiedDate: time.Now(),
		Status:           t.Status,
		CreatedBy:        t.CreatedBy,
		ProjectID:        t.ProjectID,
		Estimate:         t.Estimate,
		Remaining:        t.Remaining,
		SprintID:         t.SprintID,
	}

	if len(newRecord.OwnerName) == 0 {
		newRecord.OwnerName = "Unassigned"
	}
	return newRecord, nil
}

//UpdateTask update test record
func UpdateTask(t *Task) (*Task, error) {
	if err := validateNewTask(t); err != nil {
		return nil, err
	}
	t.LastModifiedDate = time.Now()
	return t, nil
}

func validateNewTask(t *Task) error {

	if len(t.TaskName) == 0 {
		return errors.New(NameRequired)
	}

	if len(t.OwnerID) == 0 {
		return errors.New(OwnerIDRequired)
	}

	if len(t.Status) == 0 {
		return errors.New(StatusRequired)
	}

	if len(t.ProjectID) == 0 {
		return errors.New(ProjectIDRequired)
	}
	if len(t.SprintID) == 0 {
		return errors.New(SprintIDRequired)
	}

	if t.Status == "Closed" && t.Remaining > 0 {
		return errors.New(RemainingHoursMustBeEqaulToZero)
	}
	return nil
}
