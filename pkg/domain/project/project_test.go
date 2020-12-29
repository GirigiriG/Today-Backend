package project_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/assert"

	"github.com/GirigiriG/Clean-Architecture-golang/tools"

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
	assert.Equal(t, project.PROJECT_NAME_REQUIRED, err.Error())
}

func Test_Status_Required(t *testing.T) {
	newProjectRecord := createNewProjectTestData()
	newProjectRecord.Status = ""

	_, err := project.NewProject(newProjectRecord)

	assert.NotNil(t, err)
	assert.Equal(t, project.SPRINT_ID_REQUIRED, err.Error())
}

func Test_Sprint_ID_Required(t *testing.T) {
	newProjectRecord := createNewProjectTestData()
	newProjectRecord.SprintID = uuid.Nil

	_, err := project.NewProject(newProjectRecord)
	assert.NotNil(t, err)
	assert.Equal(t, project.SPRINT_ID_REQUIRED, err.Error())
}

func Test_Project_Percentage_Must_Equal_100(t *testing.T) {
	newProjectRecord := createNewProjectTestData()
	newProjectRecord.Status = "Completed"
	newProjectRecord.PercentComplete = 99
	_, err := project.NewProject(newProjectRecord)

	assert.NotNil(t, err)
	assert.Equal(t, project.PROJECT_PERCENT_NOT_100, err.Error())
}

func createNewProjectTestData() *project.Project {
	id := tools.CreateUUID()
	sprintRecordID := tools.CreateUUID()
	createdBy := tools.CreateUUID()

	return &project.Project{
		ID:               id,
		Status:           "Active",
		SprintID:         sprintRecordID,
		ProjectName:      "test project name",
		Description:      "project descriptions",
		CreatredBy:       createdBy,
		Created_Date:     time.Now(),
		PercentComplete:  100,
		LastModifiedDate: time.Now(),
	}
}
