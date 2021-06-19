package application

import (
	"log"
	"net/http"

	controller "github.com/shoooooman/sample-ddd-app/application/controller/user"
)

func Run() {
	router := http.NewServeMux()
	router.HandleFunc("/users", controller.UserHandler)

	http.ListenAndServe(":8080", logger(router))
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := r.RemoteAddr
		method := r.Method
		path := r.URL.Path
		log.Printf("%s [%s] %s", addr, method, path)
		h.ServeHTTP(w, r)
	})
}
