package controller

import (
	"net/http"
	"strconv"

	"github.com/shoooooman/sample-ddd-app/application/response"
)

type userGetRequest struct {
	UserID int
}

type userGetResponse struct {
	UserID int    `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.FormValue("user_id")
	if userIDStr == "" {
		response.BadRequest(w, "user_id is required")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.BadRequest(w, "user_id is not valid")
		return
	}

	request := &userGetRequest{
		UserID: userID,
	}

	user := userUsecase.FindUser(request.UserID)
	if user == nil {
		err := response.OK(w, &userGetResponse{})
		if err != nil {
			response.InternalServerError(w, err.Error())
		}
		return
	}

	resp := &userGetResponse{
		UserID: user.ID,
		Name:   user.Name,
	}

	if err = response.OK(w, resp); err != nil {
		response.InternalServerError(w, err.Error())
	}
}
