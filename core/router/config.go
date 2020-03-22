package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init initializes the core service's router.
func Init(host, port string) {
	m := mux.NewRouter()
	configureRoutes(m)

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Fatal(http.ListenAndServe(addr, m))
}
