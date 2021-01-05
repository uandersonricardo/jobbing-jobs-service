package helpers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, message string) {
	err := Error{message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	res, _ := json.Marshal(err)

	w.Write([]byte(res))

	return
}
