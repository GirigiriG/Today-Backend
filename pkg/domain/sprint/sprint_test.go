package sprint_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

	"github.com/stretchr/testify/assert"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/domain/sprint"
)

func TestNewSprintCreationSuccess(t *testing.T) {
	s := testData()
	ID := tools.GenerateStringUUID()
	_, err := sprint.NewSprint(s, ID)

	assert.Equal(t, nil, err)
}

func TestNewSprintCreationFailure(t *testing.T) {
	s := testData()
	s.SprintName = ""

	ID := tools.GenerateStringUUID()
	_, err := sprint.NewSprint(s, ID)

	assert.Equal(t, sprint.ErrSprintNameIsRquired, err.Error())
}

func TestNewSprintStartDateRequired(t *testing.T) {
	s := &sprint.Sprint{}
	ID := tools.GenerateStringUUID()
	s.StartDate = time.Now().Add(67 * 60)

	fmt.Println(s.StartDate.IsZero())
	_, err := sprint.NewSprint(s, ID)

	assert.Equal(t, sprint.ErrSprintNameIsRquired, err.Error())
}

func TestNewSprintStartDateGreaterOrEqualNow(t *testing.T) {
	s := &sprint.Sprint{}

	s.SprintName = "test"
	s.StartDate = time.Now().Add(36 * time.Hour)
	s.EndDate = time.Now().Add(12 * time.Hour)
	ID := tools.GenerateStringUUID()

	_, err := sprint.NewSprint(s, ID)
	assert.Equal(t, sprint.ErrStartDateMustBeTodayOrGreater, err.Error())
}

func TestNewSprintEndDateRequired(t *testing.T) {
	s := testData()
	s.EndDate = time.Time{}

	ID := tools.GenerateStringUUID()
	_, err := sprint.NewSprint(s, ID)
	assert.Equal(t, sprint.ErrEndDatedIsRquired, err.Error())
}

func TestNewSprintEndDateAfterStartDate(t *testing.T) {
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
