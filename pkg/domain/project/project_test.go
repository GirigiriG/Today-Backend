package project_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/project"
)

func Test_Create_Project_Record_Successfully(t *testing.T) {
	newProjectRecord := createNewProjectTestData()

	_, err := project.NewProject(newProjectRecord)
	assert.NoError(t, err)
}

func Test_Project_Name_Required(t *testing.T) {
	newProjectRecord := createNewProjectTestData()
	newProjectRecord.ProjectName = ""

	_, err := project.NewProject(newProjectRecord)

	assert.NotNil(t, err)
	assert.Equal(t, project.ProjectNameRequired, err.Error())
}

func Test_Status_Required(t *testing.T) {
	newProjectRecord := createNewProjectTestData()
	newProjectRecord.Status = ""

	_, err := project.NewProject(newProjectRecord)

	assert.NotNil(t, err)
	assert.Equal(t, project.StatusRequired, err.Error())
}

func Test_Sprint_ID_Required(t *testing.T) {
	newProjectRecord := createNewProjectTestData()
	newProjectRecord.SprintID = ""

	_, err := project.NewProject(newProjectRecord)
	assert.NotNil(t, err)
	assert.Equal(t, project.SprintIDRequired, err.Error())
}

func Test_Project_Percentage_Must_Equal_100(t *testing.T) {
	newProjectRecord := createNewProjectTestData()
	newProjectRecord.Status = "Completed"
	newProjectRecord.PercentComplete = 99
	_, err := project.NewProject(newProjectRecord)

	assert.NotNil(t, err)
	assert.Equal(t, project.PercentageMustEqaul100, err.Error())
}

func createNewProjectTestData() *project.Project {
	Id := tools.GenerateStringUUID()
	sprintRecordID := tools.GenerateStringUUID()
	createdBy := tools.GenerateStringUUID()

	return &project.Project{
		ID:               Id,
		Status:           "Active",
		SprintID:         sprintRecordID,
		ProjectName:      "test project name",
		Description:      "project descriptions",
		CreatedBy:        createdBy,
		CreatedDate:      time.Now(),
		PercentComplete:  100,
		LastModifiedDate: time.Now(),
	}
}
