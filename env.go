package main

import (
	"log"

	"github.com/navaz-alani/dotenv"
)

var env *dotenv.Env

// required specifies which keys MUST be read
// from the environment variables. Failure to
// read these variables will cause the program
// to exit with status 1.
var required = []string{
	"APP",
	"HOST",
	"PORT",
	"DB_URI",
}

// envInit initializes some application parameters
// from the environment variables.
func envInit(src string) {
	e, err := dotenv.Load(src, true)
	if err != nil {
		log.Println(err)
		log.Fatalln("err: unable to initialize environment")
	}
	undef := e.CheckRequired(required)
	if undef != nil {
		log.Fatalln("err: expected definition in environment", undef)
	}
	env = e
}
