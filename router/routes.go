package router

import (
	"log"
	"net/http"

	"github.com/navaz-alani/go-talk/core"
)

func configureRoutes(mux *http.ServeMux) {
	userCreateMW, err := core.EMux.CreationMiddleware("user")
	if err != nil {
		log.Println("mw setup fail")
		log.Fatal(err)
	}
	registerUser := userCreateMW(http.HandlerFunc(core.RegisterUser))

	mux.HandleFunc("/register/user", registerUser.ServeHTTP)
}
