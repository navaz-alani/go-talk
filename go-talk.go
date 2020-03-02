package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/navaz-alani/go-talk/core"
	"github.com/navaz-alani/go-talk/db"
	"github.com/navaz-alani/go-talk/env"
	"github.com/navaz-alani/go-talk/router"
)

var (
	appName string
	dbUri   string
	port    string
)

/*
main function configures the services required
for the application and initializes them with
respect to any dependencies.
*/
func main() {
	// initialize environment vars
	env.Init(".env")
	// check if required params are specified
	env.VerifyRequired()

	// read params for app init
	appName = env.Get("APP")
	dbUri = env.Get("DB_URI")
	port = env.Get("PORT")

	// initialize database client
	disconnect := db.Init(dbUri)
	defer func() {
		_ = disconnect(context.TODO())
	}()

	// initialize messaging multiplexer
	core.Init()

	// initialize backend API
	app := router.Init()
	addr := fmt.Sprintf(":%s", port)

	log.Printf("%s: listening on %s\n", appName, addr)
	_ = http.ListenAndServe(addr, app)
}
