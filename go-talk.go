package main

import (
	"context"
	"log"

	"github.com/navaz-alani/go-talk/core"
	"github.com/navaz-alani/go-talk/db"
)

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
