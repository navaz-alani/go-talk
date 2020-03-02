package client

import (
	"context"

	"github.com/navaz-alani/go-talk/core/comm"
)

/*
Client is an interface describing the capabilities
of a client to the messaging system.
*/
type Client interface {
	/*
		Ping reports whether the client is
		currently active.
	*/
	Ping(ctx context.Context) (t comm.OPToken)
	/*
		Send attempts to send the given payload
		to the client. OPToken returned can reports
		outcome. OK if t.Error() == nil.
	*/
	Send(ctx context.Context, payload comm.Payload) (t comm.OPToken)
	/*
		Listen returns a channel which over which
		incoming payloads can be received.
		This channel is buffered and the buffer stats
		can be obtained by calling the InBuffer() method.
	*/
	Listen() (in <-chan comm.Payload, quit <-chan bool)
	/*
		InBuffer reports the incoming channel's buffer
		status. If buff.Cap <= buff.Curr, the buffer is
		full and all incoming messages will fail to
		send (on the client side).
	*/
	InBuffer() (buff comm.BuffStat)
	/*
		OutBuffer reports the outgoing channel's buffer
		status. If buff.Cap <= buff.Curr, the buffer is
		full and all outgoing messages will be blocked
		from sending.
	*/
	OutBuffer() (buff comm.BuffStat)
	/*
		Quit ends the connection between the client and
		the service. If the connection closed gracefully,
		t.Error() will be nil.
		After Quit has been called successfully, the client
		is no longer considered active and all messages to
		this client will be stored in the database for
		distribution upon the next connect.
	*/
	Quit() (t comm.OPToken)
}
