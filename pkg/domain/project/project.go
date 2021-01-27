package project

import (
	"errors"
	"time"
)

//ErrProjectNameRequired : project name rquired
var ErrProjectNameRequired = errors.New("Please provide project name")

//ErrSprintIDRequired : sprint id not provided
var ErrSprintIDRequired = errors.New("Please provide sprint id")

//ErrStatusRequired : status not provided
var ErrStatusRequired = errors.New("Please provide status")

//ErrPercentageMustEqaul100 :  percent conmplete must equal 100
var ErrPercentageMustEqaul100 = errors.New("Project percentage does not equal 100 please make sure all task are complete")

//Project : project entity
type Project struct {
	ID               string
	Status           string `json:"status"`
	SprintID         string
	ProjectName      string
	CreatedBy        string
	CreatedDate      time.Time
	LastModifiedDate time.Time
	Description      string
	PercentComplete  int
}

//NewProject create a new project record
func NewProject(p *Project) (*Project, error) {
	if err := validateProjectRecord(p); err != nil {
		return nil, err
	}

	return &Project{
		ID:               p.ID,
		Status:           p.Status,
		SprintID:         p.SprintID,
		ProjectName:      p.ProjectName,
		CreatedDate:      time.Now(),
		CreatedBy:        p.CreatedBy,
		LastModifiedDate: time.Now(),
		Description:      p.Description,
		PercentComplete:  p.PercentComplete,
	}, nil
}

//UpdateProject : update project record
func UpdateProject(p *Project) (*Project, error) {
	if err := validateProjectRecord(p); err != nil {
		return nil, err
	}

	p.LastModifiedDate = time.Now()
	return p, nil
}

func validateProjectRecord(p *Project) error {
	if len(p.ProjectName) == 0 {
		return ErrProjectNameRequired
	}
	if len(p.Status) == 0 {
		return ErrStatusRequired
	}

	if len(p.SprintID) == 0 {
		return ErrSprintIDRequired
	}
	if p.Status == "Completed" && p.PercentComplete != 100 {
		return ErrPercentageMustEqaul100
	}
	return nil
}
