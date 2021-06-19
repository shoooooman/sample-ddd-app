package application

import (
	"net/http"

	controller "github.com/shoooooman/sample-ddd-app/application/controller/user"
)

func Run() {
	http.HandleFunc("/users", controller.UserHandler)

	http.ListenAndServe(":8080", nil)
}
