// services/class_service.go
package services

import (
	"errors"

	"api/models"
)

type ClassService struct {
	Classes []models.Class
}

func NewClassService() *ClassService {
	return &ClassService{
		Classes: []models.Class{},
	}
}

func (s *ClassService) CreateClass(class models.Class) error {
	if class.StartDate.After(class.EndDate) {
		return errors.New("Start date cannot be after end date")
	}

	s.Classes = append(s.Classes, class)
	return nil
}
