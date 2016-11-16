package controllers

import (
	"net/http"
	"todoisAPI/dao"
	"strconv"
	"github.com/gorilla/mux"
	"todoisAPI/services"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err:= strconv.Atoi(mux.Vars(r)["id"])
	if(err != nil){
		services.ResponseError(w,http.StatusBadRequest,"Error en el valor del parametro id", err.Error())
		return
	}
	user, err :=  dao.GetUserById(id)
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
