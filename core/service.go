package core

import (
	"github.com/navaz-alani/entity/multiplexer"

	"github.com/navaz-alani/go-talk/core/mux"
)

var (
	EMux *multiplexer.EMux
	CMux *mux.CMux
)

/*
Init initializes the application's core service
for client connection and interaction handling.
*/
func Init() {
	CMux = mux.New()
	CMux.Handle()
}
