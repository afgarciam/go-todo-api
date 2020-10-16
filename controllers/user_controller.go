package controllers

import (
	"net/http"
	"go-todo-apidao"
	"strconv"
	"github.com/gorilla/mux"
	"go-todo-apiservices"
)

type UserController struct {}

var userDao *dao.UserDAO

func (ctrl *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	id, err:= strconv.Atoi(mux.Vars(r)["id"])
	if(err != nil){
		services.ResponseError(w,http.StatusBadRequest,"Error en el valor del parametro id", err.Error())
		return
	}
	user, err :=  userDao.GetById(id)
	if(err != nil){
		services.ResponseError(w,http.StatusInternalServerError,"Error al obtener los datos del usuario", err.Error())
		return
	}
	if(user.ID <= 0){
		services.ResponseError(w,http.StatusNotFound,"No se encontro el usuario indicado",err.Error())
		return
	}
	services.ResponseData(w, user)
}
