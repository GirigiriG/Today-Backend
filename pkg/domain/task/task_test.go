package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/task"
)

func TestCreateTaskRecordSucessfully(t *testing.T) {
	newTask := createTaskRecord()
	_, err := task.NewTask(newTask)
	assert.Equal(t, nil, err)
}

func TestNameRequired(t *testing.T) {
	newTask := createTaskRecord()
	newTask.TaskName = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.NameRequired)
}

func TestOwnerIDRequired(t *testing.T) {
	newTask := createTaskRecord()
	newTask.OwnerID = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.OwnerIDRequired)
}

func TestStatusRequired(t *testing.T) {
	newTask := createTaskRecord()
	newTask.Status = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.StatusRequired)
}

func TestProjectIdRequired(t *testing.T) {
	newTask := createTaskRecord()
	newTask.ProjectID = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.ProjectIDRequired)
}

func TestSprintIdRequired(t *testing.T) {
	newTask := createTaskRecord()
	newTask.SprintID = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.SprintIDRequired)
}

func TestContainsRemainigHours(t *testing.T) {
	newTask := createTaskRecord()
	newTask.Remaining = 3

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.RemainingHoursMustBeEqaulToZero)
}

func createTaskRecord() *task.Task {
	return &task.Task{
		TaskName:  "save the day",
		OwnerID:   "12345",
		Status:    "Closed",
		CreatedBy: "ME",
		ProjectID: "4477",
		Estimate:  77,
		Remaining: 0,
		SprintID:  "7744",
	}

}
