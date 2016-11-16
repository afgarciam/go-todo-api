package controllers

import (
	"net/http"
	"encoding/json"
	"todoisAPI/models"
	"todoisAPI/services"
	"todoisAPI/dao"
)

type AuthController struct {}

var authDao *dao.UserDAO

func (ctrl *AuthController) Login(w http.ResponseWriter, r *http.Request)  {
	loginData := models.LoginData{}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if(err != nil){
		services.ResponseError(w,http.StatusBadRequest,"Error la estructura para el login no es correcta", err.Error())
		return
	}

	u, err := userDao.Validate(&loginData)
	if(err != nil){
		services.ResponseError(w,http.StatusInternalServerError,"Error al validar la credenciales del usuario", err.Error())
		return
	}
	if(u.ID <= 0){
		services.ResponseError(w,http.StatusUnauthorized,"Error, credenciales incorrectas o usuario inactivo", err.Error())
		return
	}

	token, err := services.GenerateToken(u.Email)
	if(err != nil){
		services.ResponseError(w,http.StatusInternalServerError,"Error al generar el token de acceso", err.Error())
		return
	}

	services.ResponseData(w, token)
}
