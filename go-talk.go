/*
                  _        _ _
  __ _  ___       | |_ __ _| | | __
 / _` |/ _ \ _____| __/ _` | | |/ /
| (_| | (_) |_____| || (_| | |   <
 \__, |\___/       \__\__,_|_|_|\_\
  |___/

© 2020 Navaz Alani
Contact : nalani@uwaterloo.ca
*/

package main

import (
	"context"
	"log"

	"github.com/navaz-alani/go-talk/core"
	"github.com/navaz-alani/go-talk/db"
)

// main reads parameters from the environment variables
// and injects them into the application's services.
func main() {
	envInit(".env")

	disconnect := db.Init(env.Get("DB_URI"))
	defer func() {
		err := disconnect(context.Background())
		if err != nil {
			log.Println("err: failed to disconnect db client")
		}
	}()

	core.Init(env.Get("HOST"), env.Get("PORT"), db.Db())
}
