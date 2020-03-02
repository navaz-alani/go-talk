package router

import (
	"net/http"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	configureRoutes(mux)

	return mux
}
