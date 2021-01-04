package sprint

import (
	"errors"
	"fmt"
	"time"
)

const (
	//SprintNameIsRquired : sprint name required
	SprintNameIsRquired = "Please provide sprint name"
	//StartDatedIsRquired : start date required
	StartDatedIsRquired = "Please provide start date"
	//StartDateMustBeTodayOrGreater : start date more be today or greater
	StartDateMustBeTodayOrGreater = "Start date must be greater than today"
	//EndDatedIsRquired : End date must be greater than start date
	EndDatedIsRquired = "Please provide end date"
	//EndDateIsGreaterThanStartDate : End date must be greater than start date
	EndDateIsGreaterThanStartDate = "End date must be greater than start date"
)

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
		return errors.New(SprintNameIsRquired)
	}

	fmt.Println(s.StartDate.IsZero())
	if s.StartDate.IsZero() {
		return errors.New(StartDatedIsRquired)
	}

	if s.EndDate.IsZero() {
		return errors.New(EndDatedIsRquired)
	}

	if !s.StartDate.Before(time.Now().Add(0 * time.Minute)) {
		return errors.New(StartDateMustBeTodayOrGreater)
	}

	if !s.StartDate.Before(s.EndDate) {
		return errors.New(EndDateIsGreaterThanStartDate)
	}

	return nil
}
