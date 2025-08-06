package resps

import (
	"encoding/json"
	"log"
	"net/http"
)

// return a resp in json
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

func ERROR(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Error string `json:"erro"`
	}{
		Error: erro.Error(),
	})
}

func JSONpretty(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	jsonData, erro := json.MarshalIndent(data, "", "    ")
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write(jsonData)
}
