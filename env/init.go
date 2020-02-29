package env

import (
	"log"

	golangdotenv "github.com/navaz-alani/golang-dotenv"
)

var env golangdotenv.Env

/*
Init initializes the application's environment
variables from the given source file.
*/
func Init(source string) {
	e, err := golangdotenv.Load(source)
	if err != nil {
		log.Println("env: error - env load fail")
		log.Fatal(err)
	}

	env = e
}

/*
Get retrieves the value of the given get from the
application's environment variables.
*/
func Get(key string) string {
	return env.Get(key)
}
