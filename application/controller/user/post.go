package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type userPostRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

type userPostResponse struct {
	Message string `json:"message"`
}

func handlePostUser(w http.ResponseWriter, r *http.Request) {
	// TODO: POSTの共通の処理をいい感じにまとめる
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var request userPostRequest
	json.Unmarshal(body, &request)

	_, err = userUsecase.CreateUser(request.UserID, request.Name)
	// TODO: Responseの作り方まとめる
	if err != nil {
		response := &userPostResponse{
			Message: err.Error(),
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			return
		}

		w.Write(responseJSON)
		return
	}

	response := &userPostResponse{
		Message: "Created!",
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(responseJSON)
}
