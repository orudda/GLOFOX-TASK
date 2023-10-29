// tests/class_test.go
package tests

import (
	"api/models"
	"api/services"
	"api/utils"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClass(t *testing.T) {
	class := models.Class{
		Name:      "Test Class",
		StartDate: utils.CustomTime{Time: time.Now()},
		EndDate:   utils.CustomTime{Time: time.Now().Add(time.Hour * 24 * 7)},
		Capacity:  10,
	}

	assert.Equal(t, "Test Class", class.Name)
	assert.True(t, class.StartDate.Before(class.EndDate.Time))
	assert.Equal(t, 10, class.Capacity)
}

func TestCreateClass(t *testing.T) {
	classService := services.DBClassService()
	class := models.Class{
		Name:      "Test Class",
		StartDate: utils.CustomTime{Time: time.Now()},
		EndDate:   utils.CustomTime{Time: time.Now().Add(time.Hour * 24 * 7)},
		Capacity:  10,
	}

	err := classService.CreateClass(class)
	assert.NoError(t, err)
	assert.Len(t, classService.GetClasses(), 8)
}

func TestGetClassByID(t *testing.T) {

	classService := services.DBClassService()
	class := models.Class{
		Name:      "Test Class",
		StartDate: utils.CustomTime{Time: time.Now()},
		EndDate:   utils.CustomTime{Time: time.Now().Add(time.Hour * 24 * 7)},
		Capacity:  10,
	}

	err := classService.CreateClass(class)
	// Test case 1: Valid class ID

	validID := 1
	classes, err := classService.GetClassByID(validID)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if classes.ID != validID {
		t.Errorf("Expected class ID %d, got: %d", validID, class.ID)
	}
}
func TestGetClassByIDWithNoValidId(t *testing.T) {

	classService := services.DBClassService()

	// Test case 2: Invalid class ID
	invalidID := 3
	_, err := classService.GetClassByID(invalidID)
	if err == nil {
		t.Error("Expected an error, but got none")
	}
	if errors.Is(err, errors.New("Class not found")) {
		t.Errorf("Expected 'Class not found' error, got: %v", err)
	}
}

func TestClassNotAvailable(t *testing.T) {
	classService := services.DBClassService()

	err := classService.IsClassAvailable(utils.CustomTime{Time: time.Now()})

	if err != nil {
		t.Error("Expected an error, but got none")
	}
}

func TestDecrementCapacityUnavailable(t *testing.T) {
	classService := services.DBClassService()

	err := classService.DecrementClassCapacity(0)

	if err.Error() != "Class not available on the requested date" {
		t.Errorf("expected error message, but got none")
	}
}

func TestIncrementCapacityUnavailable(t *testing.T) {
	classService := services.DBClassService()

	err := classService.IncrementClassCapacity(100)

	if err.Error() != "Class not available on the requested date to increment capacity" {
		t.Errorf("expected error message, but got none")
	}
}
func TestDeleteClassNotFound(t *testing.T) {
	classService := services.DBClassService()

	err := classService.DeleteClass(0)

	if err.Error() != "Class not found" {
		t.Errorf("expected error message, but got none")
	}
}
func TestUpdateClassNotFound(t *testing.T) {
	classService := services.DBClassService()
	class := models.Class{
		Name:      "Test Class",
		StartDate: utils.CustomTime{Time: time.Now()},
		EndDate:   utils.CustomTime{Time: time.Now().Add(time.Hour * 24 * 7)},
		Capacity:  10,
	}

	err := classService.UpdateClass(class)

	if err.Error() != "Class not found" {
		t.Errorf("expected error message, but got none")
	}
}
