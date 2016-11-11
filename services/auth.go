package services

import (
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"context"
)
type ApiClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func VerifiyToken(w http.ResponseWriter, r * http.Request, next http.HandlerFunc)  {
	tokenHeader := r.Header.Get("Authorization")
	if(tokenHeader == ""){
		ResponseError(w,http.StatusForbidden,"Error el token de acceso no enviado", "")
		return
	}
	tokenString := ""
	if(strings.ToLower(tokenHeader[0:6]) == "bearer"){
		tokenString = tokenHeader[7:]
	}
	token, err := jwt.ParseWithClaims(tokenString, &ApiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("_millavesupersecreta_"), nil
	})
	if(!token.Valid){
		ResponseError(w,http.StatusForbidden,"Error el token de acceso no es valido", err.Error())
		return
	}
	claims, _ :=token.Claims.(*ApiClaims)
	ctx := context.WithValue(r.Context(), "user_email", claims.Email)
	next(w,r.WithContext(ctx))
}

func GenerateToken(userEmail string) (interface{}, error) {
	signing := []byte("_millavesupersecreta_")


	type Token struct {
		Token string `json:"token"`
	}
	ts := Token{}

	claims := ApiClaims{
		userEmail,
		jwt.StandardClaims{
			ExpiresAt: 150000000,
			Issuer:    "agarcia588",
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



