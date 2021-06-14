package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type userGetRequest struct {
	UserID int
}

type userGetResponse struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	// TODO: GETの共通の処理をいい感じにまとめる
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		log.Fatal(err)
	}

	request := &userGetRequest{
		UserID: userID,
	}

	user := userUsecase.FindUser(request.UserID)
	if user == nil {
		w.Write([]byte{})
		return
	}

	response := &userGetResponse{
		UserID: user.ID,
		Name:   user.Name,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(responseJSON)
}
