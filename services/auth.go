package services

import (
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"context"
	"time"
)

type apiClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type token struct {
	Token string `json:"token"`
}

func VerifyToken(w http.ResponseWriter, r * http.Request, next http.HandlerFunc)  {
	tokenHeader := r.Header.Get("Authorization")
	if(tokenHeader == ""){
		ResponseError(w,http.StatusForbidden,"Error el token de acceso no enviado", "")
		return
	}
	tokenString := ""
	if(strings.ToLower(tokenHeader[0:6]) == "bearer"){
		tokenString = tokenHeader[7:]
	}
	token, err := jwt.ParseWithClaims(tokenString, &apiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("_millavesupersecreta_"), nil
	})
	if(!token.Valid){
		ResponseError(w,http.StatusForbidden,"Error el token de acceso no es valido", err.Error())
		return
	}
	claims, _ :=token.Claims.(*apiClaims)
	ctx := context.WithValue(r.Context(), "user_email", claims.Email)
	next(w,r.WithContext(ctx))
}

func GenerateToken(userEmail string) (interface{}, error) {
	signing := []byte("_millavesupersecreta_")
	ts := token{}
	claims := apiClaims{
		userEmail,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "Todo is Api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString(signing)
	if(err != nil){
		return  ts, err
	}
	ts.Token = tokenSigned

	return ts, nil
}




