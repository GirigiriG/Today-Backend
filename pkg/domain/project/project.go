package project

import (
	"errors"
	"time"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"
)

const (
	ProjectNameRequired    = "Please provide project name"
	SprintIDRequired       = "Please provide sprint id"
	StatusRequired         = "Please provide status"
	PercentageMustEqaul100 = "Project percentage does not equal 100 please make sure all task are complete"
)

type Project struct {
	ID               string
	Status           string
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
	pid := tools.CreateUUID()

	return &Project{
		ID:               pid,
		Status:           p.Status,
		SprintID:         p.SprintID,
		ProjectName:      p.ProjectName,
		CreatedDate:      p.CreatedDate,
		CreatedBy:        p.CreatedBy,
		LastModifiedDate: p.LastModifiedDate,
		Description:      p.Description,
		PercentComplete:  p.PercentComplete,
	}, nil
}

func validateProjectRecord(p *Project) error {
	if len(p.ProjectName) == 0 {
		return errors.New(ProjectNameRequired)
	}
	if len(p.Status) == 0 {
		return errors.New(StatusRequired)
	}

	if len(p.SprintID) == 0 {
		return errors.New(SprintIDRequired)
	}
	if p.Status == "Completed" && p.PercentComplete != 100 {
		return errors.New(PercentageMustEqaul100)
	}
	return nil
}
