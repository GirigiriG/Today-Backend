package sprint_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

	"github.com/stretchr/testify/assert"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/sprint"
)

func Test_NewSprint_Creation_Success(t *testing.T) {
	s := testData()
	ID := tools.GenerateStringUUID()
	_, err := sprint.NewSprint(s, ID)

	assert.Equal(t, nil, err)
}

func Test_NewSprint_Creation_Failure(t *testing.T) {
	s := testData()
	s.SprintName = ""

	ID := tools.GenerateStringUUID()
	_, err := sprint.NewSprint(s, ID)

	assert.Equal(t, sprint.SprintNameIsRquired, err.Error())
}

func Test_NewSprint_StartDate_Required(t *testing.T) {
	s := &sprint.Sprint{}
	ID := tools.GenerateStringUUID()
	s.StartDate = time.Now().Add(67 * 60)

	fmt.Println(s.StartDate.IsZero())
	_, err := sprint.NewSprint(s, ID)

	assert.Equal(t, sprint.SprintNameIsRquired, err.Error())
}

func Test_NewSprint_StartDate_Greater_Or_Equal_Now(t *testing.T) {
	s := &sprint.Sprint{}

	s.SprintName = "test"
	s.StartDate = time.Now().Add(36 * time.Hour)
	s.EndDate = time.Now().Add(12 * time.Hour)
	ID := tools.GenerateStringUUID()

	_, err := sprint.NewSprint(s, ID)
	assert.Equal(t, sprint.StartDateMustBeTodayOrGreater, err.Error())
}

func Test_NewSprint_EndDate_Required(t *testing.T) {
	s := testData()
	s.EndDate = time.Time{}

	ID := tools.GenerateStringUUID()
	_, err := sprint.NewSprint(s, ID)
	assert.Equal(t, sprint.EndDatedIsRquired, err.Error())
}

func Test_NewSprint_EndDate_After_StartDate(t *testing.T) {
	s := testData()

	ID := tools.GenerateStringUUID()
	_, err := sprint.NewSprint(s, ID)
	assert.Equal(t, nil, err)
}

func testData() *sprint.Sprint {
	return &sprint.Sprint{
		SprintName: "test sprint",
		StartDate:  time.Now(),
		EndDate:    time.Now().Add(36 * time.Hour),
	}
}