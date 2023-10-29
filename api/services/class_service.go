// services/class_service.go
package services

import (
	"errors"

	"api/models"
	"api/utils"
)

type ClassService struct {
	Classes []models.Class
	nextID  int
}

func DBClassService() *ClassService {
	return &ClassService{
		Classes: []models.Class{},
		nextID:  1, //This is used to assign unique identifiers to new classes
	}
}

func (s *ClassService) CreateClass(class models.Class) error {
	if class.StartDate.Time.After(class.EndDate.Time) {
		return errors.New("Start date cannot be after end date")
	}

	currentDate := class.StartDate.Time // Access the time.Time value

	for !currentDate.After(class.EndDate.Time) { // Continue until currentDate is after the end date
		// Create a new class instance with the same details, except for the StartDate and EndDate
		newClass := models.Class{
			Name:      class.Name,
			StartDate: utils.CustomTime{Time: currentDate},
			EndDate:   utils.CustomTime{Time: currentDate},
			Capacity:  class.Capacity,
		}

		newClass.ID = s.nextID
		s.Classes = append(s.Classes, newClass)
		s.nextID++

		// Increment the currentDate by one day
		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return nil
}

func (s *ClassService) GetClasses() []models.Class {
	return s.Classes
}

func (s *ClassService) GetClassByID(id int) (models.Class, error) {
	for _, c := range s.Classes {
		if c.ID == id {
			return c, nil
		}
	}
	return models.Class{}, errors.New("Class not found")
}
