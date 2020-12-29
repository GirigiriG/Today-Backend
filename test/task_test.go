package task_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/task"
)

func Test_CreateTaskRecord_Sucessfully(t *testing.T) {
	newTask := createTaskRecord()
	_, err := task.NewTask(newTask)
	if err != nil {
		t.Fatalf("Name not provided")
	}
}

func Test_Task_Name_Required(t *testing.T) {
	newTask := createTaskRecord()
	newTask.TaskName = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.NAME_RQUIRED)
}

func Test__Task_Owner_ID_Required(t *testing.T) {
	newTask := createTaskRecord()
	newTask.OwnerID = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.OWNER_ID_REQUIRED)
}

func Test__Task_Status_Required(t *testing.T) {
	newTask := createTaskRecord()
	newTask.Status = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.STATUS_REQUIRED)
}

func Test__Task_Project_Id_Required(t *testing.T) {
	newTask := createTaskRecord()
	newTask.ProjectID = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.PROJECT_ID_REQUIRED)
}

func Test__Task_Sprint_Id_Required(t *testing.T) {
	newTask := createTaskRecord()
	newTask.SprintID = ""

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.SPRINT_ID_REQUIRED)
}

func Test__Task_Contains_Remainig_Hours(t *testing.T) {
	newTask := createTaskRecord()
	newTask.Remaining = 3

	_, err := task.NewTask(newTask)

	assert.Equal(t, err.Error(), task.HAS_REMAINING_HOURS)
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
