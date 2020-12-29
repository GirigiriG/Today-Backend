package project

import (
	"errors"
	"time"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"

	uuid "github.com/satori/go.uuid"
)

const (
	PROJECT_NAME_REQUIRED   = "Please provide project name"
	SPRINT_ID_REQUIRED      = "Please provide sprint id"
	PROJECT_PERCENT_NOT_100 = "Project percentage does not equal 100 please make sure all task are complete"
)

type Project struct {
	ID               uuid.UUID
	Status           string
	SprintID         uuid.UUID
	ProjectName      string
	CreatredBy       uuid.UUID
	Created_Date     time.Time
	LastModifiedDate time.Time
	Description      string
	PercentComplete  int
}

func NewProject(p *Project) (*Project, error) {
	if err := validateProjectRecord(p); err != nil {
		return nil, err
	}
	return &Project{
		ID:               tools.CreateUUID(),
		Status:           p.Status,
		SprintID:         tools.CreateUUID(),
		ProjectName:      p.ProjectName,
		Created_Date:     p.Created_Date,
		CreatredBy:       p.CreatredBy,
		LastModifiedDate: p.LastModifiedDate,
		Description:      p.Description,
		PercentComplete:  p.PercentComplete,
	}, nil
}

func validateProjectRecord(p *Project) error {
	if len(p.ProjectName) == 0 {
		return errors.New(PROJECT_NAME_REQUIRED)
	}
	if len(p.Status) == 0 {
		return errors.New(SPRINT_ID_REQUIRED)
	}

	if p.SprintID == uuid.Nil {
		return errors.New(SPRINT_ID_REQUIRED)
	}
	if p.Status == "Completed" && p.PercentComplete != 100 {
		return errors.New(PROJECT_PERCENT_NOT_100)
	}
	return nil
}
