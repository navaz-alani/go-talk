package core

import (
	"log"

	"github.com/navaz-alani/entity"
	"github.com/navaz-alani/entity/multiplexer"
	"github.com/navaz-alani/entity/multiplexer/muxHandle"

	"github.com/navaz-alani/go-talk/core/chat"
	"github.com/navaz-alani/go-talk/core/router"
	"github.com/navaz-alani/go-talk/core/user"
)

// toRegister is a slice containing specifications
// of entities to register in the entity multiplexer
// for the core service.
var toRegister = []struct {
	sym    string
	def    interface{}
	handle **entity.Entity
}{
	{"user", user.User{}, &user.Users},
}

// entityInit initializes the entities used in the
// application's core service.
// Initialization involves creating an EntityMux with
// core entity definitions and initializing persistent
// storage handles.
func entityInit(db muxHandle.DBHandler) {
	var defs []interface{}
	for i := 0; i < len(toRegister); i++ {
		entry := toRegister[i]
		defs = append(defs, entry.def)
	}

	eMux, err := multiplexer.Create(db, defs...)
	if err != nil {
		log.Fatal("err: core entity init fail")
	}

	for i := 0; i < len(toRegister); i++ {
		entry := toRegister[i]
		*entry.handle = eMux.E(entry.sym)
	}
}

// Init initializes the application's core service.
func Init(host, port string, db muxHandle.DBHandler) {
	chat.Init()
	entityInit(db)
	router.Init(host, port)
}
