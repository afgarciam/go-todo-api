package services

import (
	"net/http"
	"encoding/json"
)

type Message struct {
	Code int `json:"code"`
	Message string `json:"message"`
	TechnicalMessage string `json:"technical_message"`
}

func ResponseError(w http.ResponseWriter, code int, message string, tecMessage string)   {
	m := Message{
		Code: code,
		Message: message,
		TechnicalMessage: tecMessage,
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	resJson, _ := json.Marshal(m)
	w.Write(resJson)
}

func ResponseData(w http.ResponseWriter, data interface{})  {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	resJson, _ := json.Marshal(data)
	w.Write(resJson)
}