package controllers

import (
	"net/http"
	"todoisAPI/services"
	"todoisAPI/dao"
	"todoisAPI/models"
	"encoding/json"
	//"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
)

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("user_email")

	tasks, err := dao.GetAllTasks(userEmail.(string))
	if (err != nil) {
		services.ResponseError(w, http.StatusInternalServerError, "Error al obtener las tareas del usuario", err.Error())
		return
	}
	if (len(tasks) <= 0 ) {
		services.ResponseError(w, http.StatusNotFound, "No se encontraron tareas para el usuario indicado", err.Error())
		return
	}
	services.ResponseData(w, tasks)
}


func CreateTask(w http.ResponseWriter, r *http.Request){
	userEmail := r.Context().Value("user_email")
	task := models.Task{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&task)
	if(err != nil){
		services.ResponseError(w, http.StatusBadRequest, "Error la informacion enviada no es valida", err.Error())
		return
	}
	task, err = dao.CreateTask(task, userEmail.(string))
	if (err != nil) {
		services.ResponseError(w, http.StatusInternalServerError, "Error al crear la tarea", err.Error())
		return
	}

	services.ResponseData(w,task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	id, err:= strconv.Atoi(mux.Vars(r)["id"])
	if(err != nil){
		services.ResponseError(w, http.StatusBadRequest, "Error la informacion enviada no es valida", err.Error())
		return
	}
	err = dao.DeleteTask(id)
	if (err != nil) {
		services.ResponseError(w, http.StatusInternalServerError, "Error al crear la tarea", err.Error())
		return
	}

	services.ResponseData(w,nil)
}

func UpdateTask(w http.ResponseWriter, r *http.Request){
	id, err:= strconv.Atoi(mux.Vars(r)["id"])
	if(err != nil){
		services.ResponseError(w, http.StatusBadRequest, "Error la informacion enviada no es valida", err.Error())
		return
	}
	err = dao.UpdateTask(id)
	if (err != nil) {
		services.ResponseError(w, http.StatusInternalServerError, "Error al crear la tarea", err.Error())
		return
	}

	services.ResponseData(w,nil)
}
