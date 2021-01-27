package task

import (
	"errors"
	"time"
)


	//NameRequired : Name required error messsage
	var ErrNameRequired = errors.New("Please provide a task name")
	//OwnerIDRequired : Owuner id required error messsage
	var ErrOwnerIDRequired = errors.New("Pleasae provided owner id")
	//StatusRequired : Status required error messsage
	var ErrStatusRequired = errors.New("Pleasae provided status")
	//ProjectIDRequired : Project id required error messsage
	var ErrProjectIDRequired = errors.New("Pleasae provided project id")
	//SprintIDRequired : Sprint id required error messsage
	var ErrSprintIDRequired = errors.New("Pleasae provided sprint id")
	//RemainingHoursMustBeEqaulToZero : Task must have 0 hours remaining error message
	var ErrRemainingHoursMustBeEqaulToZero = errors.New("This task still has remaining hours")


//Task Object
type Task struct {
	ID               string    `json:"id"`
	TaskName         string    `json:"name"`
	OwnerID          string    `json:"ownerId"`
	OwnerName        string    `json:"ownerName"`
	CreatedDate      time.Time `json:"createdDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
	Status           string    `json:"status"`
	CreatedBy        string    `json:"createdBy"`
	ProjectID        string    `json:"projectId"`
	Estimate         int32     `json:"estimate"`
	Remaining        int32     `json:"remaining"`
	SprintID         string    `json:"sprintId"`
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
		return ErrNameRequired
	}

	if len(t.OwnerID) == 0 {
		return ErrOwnerIDRequired
	}

	if len(t.Status) == 0 {
		return ErrStatusRequired
	}

	if len(t.ProjectID) == 0 {
		return ErrProjectIDRequired
	}
	if len(t.SprintID) == 0 {
		return ErrSprintIDRequired
	}

	if t.Status == "Closed" && t.Remaining > 0 {
		return ErrRemainingHoursMustBeEqaulToZero
	}
	return nil
}
