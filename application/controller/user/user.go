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
	// TODO: loggerをいい感じに仕込む
	log.Println(r)

	userUsecase, _ = container.DIC.Inject("UserUsecase").(usecase.UserUsecase)

	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r)
	case http.MethodPost:
		handlePostUser(w, r)
	}
}
