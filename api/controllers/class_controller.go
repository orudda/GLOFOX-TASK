// controllers/class_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/models"
	"api/services"
	"api/utils"

	"github.com/go-chi/chi"
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

func (c *ClassController) GetClasses(w http.ResponseWriter, r *http.Request) {
	classes := c.ClassService.GetClasses()
	utils.RespondWithJSON(w, http.StatusOK, classes)
}

func (c *ClassController) GetClassByID(w http.ResponseWriter, r *http.Request) {
	classID, _ := strconv.Atoi(chi.URLParam(r, "classID"))
	class, err := c.ClassService.GetClassByID(classID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, class)
}
