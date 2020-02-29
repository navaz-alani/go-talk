package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/navaz-alani/go-talk/db"
	"github.com/navaz-alani/go-talk/env"
	"github.com/navaz-alani/go-talk/router"
)

/*
main function configures the services required
for the application and initializes them with
respect to any dependencies.
*/
func main() {
	// initialize environment vars
	env.Init(".env")

	// initialize database client
	disconnect := db.Init(env.Get("DB_URI"))
	defer func() {
		_ = disconnect(context.TODO())
	}()

	// initialize backend API
	app := router.Init()
	addr := fmt.Sprintf(":%s", env.Get("PORT"))

	log.Printf("%s: listening on %s\n", env.Get("APP"), addr)
	_ = http.ListenAndServe(addr, app)
}
