// tests/class_test.go
package tests

import (
	"api/models"
	"api/services"
	"api/utils"
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
	assert.Len(t, classService.GetClasses(), 1)
}
