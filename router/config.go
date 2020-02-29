package router

import "net/http"

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	// additional config here..

	return mux
}
