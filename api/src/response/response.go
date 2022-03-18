package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return a response in JSON for a resquest
func JSON(w http.ResponseWriter, statusCode int, datas interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(datas); erro != nil {
		log.Fatal(erro)
	}
}

// Erro return a error in format JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
