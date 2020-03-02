package env

import (
	"log"
	"os"

	dotEnv "github.com/navaz-alani/golang-dotenv"
)

var (
	env dotEnv.Env
	/*
		required is a slice of strings which represent
		keys in the environment variables which must
		be defined for the application to run.
	*/
	required = []string{
		"APP",
		"PORT",
		"JWT_ISS",
		"JWT_SS",
		"DB_URI",
	}
)

/*
Init initializes the application's environment
variables from the given source file.
*/
func Init(source string) {
	e, err := dotEnv.Load(source)
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

/*
VerifyRequired verifies that the required environment variables
have been defined. If the verification fails, the program
may no longer continue execution and will exit with code 1.
*/
func VerifyRequired() {
	ok := true
	for _, key := range required {
		if key == "" {
			log.Printf("env: error - '%s' undefined", key)
			ok = false
		}
	}
	if !ok {
		os.Exit(1)
	}
}
