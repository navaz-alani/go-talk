package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Init initializes the core service's router.
func Init(host, port string) {
	m := mux.NewRouter()
	configureRoutes(m)

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, cors.AllowAll().Handler(m)))
}
