package controllers

import (
	"encoding/json"
	"go-todo-api/dao"
	"go-todo-api/models"
	"go-todo-api/services"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type TaskController struct{}

var taskDao *dao.TaskDAO

func (ctrl *TaskController) GetAll(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("user_email")

	tasks, err := taskDao.GetAll(userEmail.(string))
	if err != nil {
		services.ResponseError(w, http.StatusInternalServerError, "Error al obtener las tareas del usuario", err.Error())
		return
	}
	if len(tasks) <= 0 {
		services.ResponseError(w, http.StatusNotFound, "No se encontraron tareas para el usuario indicado", err.Error())
		return
	}
	services.ResponseData(w, tasks)
}

func (ctrl *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userEmail := r.Context().Value("user_email")
	task := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		services.ResponseError(w, http.StatusBadRequest, "Error la informacion enviada no es valida", err.Error())
		return
	}
	result, err := govalidator.ValidateStruct(task)
	if !result {
		services.ResponseError(w, http.StatusBadRequest, "Error la estructura para crear la tarea no es correcta", err.Error())
		return
	}
	task, err = taskDao.Create(task, userEmail.(string))
	if err != nil {
		services.ResponseError(w, http.StatusInternalServerError, "Error al crear la tarea", err.Error())
		return
	}

	services.ResponseData(w, task)
}

func (ctrl *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		services.ResponseError(w, http.StatusBadRequest, "Error la informacion enviada no es valida", err.Error())
		return
	}
	err = taskDao.Delete(id)
	if err != nil {
		services.ResponseError(w, http.StatusInternalServerError, "Error al crear la tarea", err.Error())
		return
	}

	services.ResponseData(w, nil)
}

func (ctrl *TaskController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		services.ResponseError(w, http.StatusBadRequest, "Error la informacion enviada no es valida", err.Error())
		return
	}
	err = taskDao.Update(id)
	if err != nil {
		services.ResponseError(w, http.StatusInternalServerError, "Error al crear la tarea", err.Error())
		return
	}

	services.ResponseData(w, nil)
}
