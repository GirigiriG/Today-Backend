package sprint

import (
	"errors"
	"fmt"
	"time"
)

//ErrSprintNameIsRquired : sprint name required
var ErrSprintNameIsRquired = errors.New("Please provide sprint name")

//ErrStartDatedIsRquired : start date required
var ErrStartDatedIsRquired = errors.New("Please provide start date")

//ErrStartDateMustBeTodayOrGreater : start date more be today or greater
var ErrStartDateMustBeTodayOrGreater = errors.New("Start date must be greater than today")

//ErrEndDatedIsRquired : End date must be greater than start date
var ErrEndDatedIsRquired = errors.New("Please provide end date")

//ErrEndDateIsGreaterThanStartDate : End date must be greater than start date
var ErrEndDateIsGreaterThanStartDate = errors.New("End date must be greater than start date")

//Sprint : A struct that holds sprint's record data
type Sprint struct {
	ID               string
	SprintName       string
	CreatedBy        string
	OwnerID          string
	StartDate        time.Time
	EndDate          time.Time
	CreatedDate      time.Time
	LastModifiedDate time.Time
}

//NewSprint : creates a new s record if record is valid
func NewSprint(s *Sprint, ID string) (*Sprint, error) {
	if err := validate(s); err != nil {
		return nil, err
	}
	s.ID = ID
	s.CreatedDate = time.Now()
	return s, nil
}

//UpdateSprint : update sprint record
func UpdateSprint(s *Sprint) (*Sprint, error) {
	if err := validate(s); err != nil {
		return nil, err
	}

	s.LastModifiedDate = time.Now()
	return s, nil
}

func validate(s *Sprint) error {
	if len(s.SprintName) == 0 {
		return ErrSprintNameIsRquired
	}

	fmt.Println(s.StartDate.IsZero())
	if s.StartDate.IsZero() {
		return ErrStartDatedIsRquired
	}

	if s.EndDate.IsZero() {
		return ErrEndDatedIsRquired
	}

	if !s.StartDate.Before(time.Now().Add(0 * time.Minute)) {
		return ErrStartDateMustBeTodayOrGreater
	}

	if !s.StartDate.Before(s.EndDate) {
		return ErrEndDateIsGreaterThanStartDate
	}

	return nil
}
