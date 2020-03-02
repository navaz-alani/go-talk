/*
Package core provides the core functionality
for client and message handling.
*/
package mux

import (
	"context"

	"github.com/navaz-alani/go-talk/core/client"
	"github.com/navaz-alani/go-talk/core/comm"
)

/*
ConnectionID is an alias for a string representing
the mux ID of a connection.
*/
type ConnectionID string

/*
Connection is a type representing a connection with
the multiplexer.
*/
type Connection struct {
	/*
		ID is a string containing the connection's
		connection ID. This is used for specifying
		payload sources and destinations.
	*/
	ID ConnectionID
	/*
		client stores a pointer to a connection to
		the peer.
	*/
	client *client.Client
	/*
		dist is a channel for payloads which
		have to be distributed.
	*/
	dist chan comm.Payload
}

/*
CMux is type which handles interactions between
message transport transmission between connected
clients.


*/
type CMux struct {
	/*
		Active is the collection of clients currently
		connected to the service.
	*/
	Active []*client.Client
}

/*
New allocates and initializes a CMux for multiplexing
connections and interactions.
*/
func New() *CMux {
	cMux := new(CMux)

	return cMux
}

/*
RegisterClient registers a websocket client to the multiplexer's
active clients.
*/
func (mux *CMux) RegisterClient(c *client.Client) {
	mux.Active = append(mux.Active, c)
	_ = (*c).Ping(context.TODO())
}

/*
Handle begins a routine for handling incoming and
outgoing messages between active clients.
*/
func (mux *CMux) Handle() {
	// todo
}

/*
QueueDist adds the payload to the distribution queue
of mux.
*/
func (mux *CMux) QueueDist() {
	// todo
}
