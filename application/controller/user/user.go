package controller

import (
	"log"
	"net/http"

	"github.com/shoooooman/sample-ddd-app/container"
	usecase "github.com/shoooooman/sample-ddd-app/usecase/user"
)

var (
	userUsecase usecase.UserUsecase
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var ok bool
	userUsecase, ok = container.DIC.Inject("UserUsecase").(usecase.UserUsecase)
	if !ok {
		log.Fatal("injection type error")
	}

	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r)
	case http.MethodPost:
		handlePostUser(w, r)
	}
}
