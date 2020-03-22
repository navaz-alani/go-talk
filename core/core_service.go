package core

import (
	"github.com/navaz-alani/entity/multiplexer/muxHandle"

	"github.com/navaz-alani/go-talk/core/router"
)

// Init initializes the application's core service.
func Init(host, port string, db muxHandle.DBHandler) {
	// todo: Initialize entities using the database
	router.Init(host, port)
}
