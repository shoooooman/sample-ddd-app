package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shoooooman/sample-ddd-app/application/response"
	usecase "github.com/shoooooman/sample-ddd-app/usecase/user"
)

type userPostRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

type userPostResponse struct {
	Message string `json:"message"`
}

func handlePostUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	// set defalut values explicitly to check null properties
	request := userPostRequest{UserID: -1, Name: ""}
	json.Unmarshal(body, &request)

	if request.UserID == -1 {
		response.BadRequest(w, "user_id is required")
		return
	}
	if request.Name == "" {
		response.BadRequest(w, "name is required")
		return
	}

	if err = userUsecase.CreateUser(request.UserID, request.Name); err != nil {
		switch e := err.(type) {
		case *usecase.ModelError:
			response.BadRequest(w, e.Error())
		case *usecase.RepositoryError:
			response.InternalServerError(w, e.Error())
		default:
			response.InternalServerError(w, e.Error())
		}
		return
	}

	resp := &userPostResponse{
		Message: "user created",
	}

	if err = response.OK(w, resp); err != nil {
		response.InternalServerError(w, err.Error())
	}
}
