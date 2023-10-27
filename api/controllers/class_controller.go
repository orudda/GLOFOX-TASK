// controllers/class_controller.go
package controllers

import (
	"encoding/json"
	"net/http"

	"api/models"
	"api/services"
	"api/utils"
)

type ClassController struct {
	ClassService *services.ClassService
}

func NewClassController(classService *services.ClassService) *ClassController {
	return &ClassController{ClassService: classService}
}

func (c *ClassController) CreateClass(w http.ResponseWriter, r *http.Request) {
	var newClass models.Class
	if err := json.NewDecoder(r.Body).Decode(&newClass); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.ClassService.CreateClass(newClass); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, newClass)
}
