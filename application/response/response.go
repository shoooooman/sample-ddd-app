package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func OK(w http.ResponseWriter, resp interface{}) error {
	respJSON, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)

	return nil
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	s := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	msgJSON, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(msgJSON)
}

func InternalServerError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	s := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	msgJSON, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(msgJSON)
}
