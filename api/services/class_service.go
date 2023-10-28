// services/class_service.go
package services

import (
	"errors"

	"api/models"
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

	class.ID = s.nextID
	s.Classes = append(s.Classes, class)
	s.nextID++
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
