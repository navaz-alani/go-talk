package client

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	ws "github.com/gorilla/websocket"

	"github.com/navaz-alani/go-talk/core/comm"
)

/*
These constants specify the payload capacities for the
input and output channels of the client.
*/
const (
	// input buffer max capacity in # of comm.Payload
	inBuffMax = 100
	// output buffer max capacity in # of comm.Payload
	outBuffMax = 100

	// ping timeout in milliseconds
	pingTimeout = 300
)

/*
WSClient is a remote client to the messaging
service.
*/
type WSClient struct {
	/*
		conn is the websocket connection established
		between a remote client and the core service.
	*/
	Conn *ws.Conn
	/*
		workers is the count of active listeners for
		this client.
	*/
	workers int
	/*
		dist is a channel which is used to push payloads
		to the multiplexer for distribution.
	*/
	dist chan<- comm.Payload
	/*
		in is the input channel for the client. It is
		buffered with a capacity specified by inBuffMax.
	*/
	in chan comm.Payload
	/*
		inCount is the number of payloads that have
		to be sent to the peer.
	*/
	inCount int
	/*
		out is the output channel for the client. It is
		buffered, with a capacity specified by outBuffMax.
	*/
	out chan comm.Payload
	/*
		outCount is the number of payloads that have to
		be sent to the multiplexer for distribution.
	*/
	outCount int
	/*
		quit is the quit
	*/
	quit chan bool
}

/*
Upgrader defines the specifications of the websocket
connection upgrader.
*/
var Upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

/*
NewWSClient allocates and initializes a new WSClient.
*/
func NewWSClient(conn *ws.Conn) WSClient {
	wsc := WSClient{
		Conn: conn,
		in:   make(chan comm.Payload, inBuffMax),
		out:  make(chan comm.Payload, outBuffMax),
		quit: make(chan bool),
	}

	return wsc
}

/*
Ping returns whether the client's websocket
connection is active.
*/
func (c WSClient) Ping(ctx context.Context) (t comm.OPToken) {
	err := c.Conn.WriteControl(ws.PingMessage, []byte("go-talk ping"),
		time.Now().Add(300*time.Millisecond))
	if err != nil {
		return err
	}

	return t
}

/*
Send sends the contents of the given payload to
the client c.
If everything goes well, the t.Error() will be nil
otherwise t.Error() will contain the encountered
error.
*/
func (c WSClient) Send(ctx context.Context, payload comm.Payload) (t comm.OPToken) {
	err := c.Conn.WriteMessage(1, payload.Raw())
	if err != nil {
		t = err
	}

	return t
}

/*
startListener starts a listener routine which
listens over the websocket connection for any
incoming payloads and forwards them to the
CMux for distribution.
*/
func (c WSClient) startListener() {
	go func() {
		defer func() {
			c.Quit()
		}()

		for {
			_, msg, err := c.Conn.ReadMessage()
			if err != nil {
				if ws.IsUnexpectedCloseError(err,
					ws.CloseGoingAway, ws.CloseAbnormalClosure) {
					close(c.in)
				}
				break
			}

			var payload comm.WSPayload
			err = json.Unmarshal(msg, &payload)
			if err != nil {
				// respond to client with error; ignore, move on
				c.Send(context.TODO(), comm.WSPayload{
					Type: 1, Str: "bad payload",
				})
				continue
			}

			if payload.Signal == "quit" {
				break
			}

			// send to multiplexer for distribution
			c.in <- payload
		}
	}()

	c.workers++
}

/*
Listen returns a buffered channel on which incoming
payloads from websocket client c are sent. It also
returns a quit channel which can be used to listen
for the quit signal.
Note that the returned channels are read-only.
*/
func (c WSClient) Listen() (in <-chan comm.Payload, quit <-chan bool) {
	if c.workers == 0 {
		c.startListener()
	}

	return c.in, c.quit
}

/*
InBuffer reports the status of the incoming payload
buffer for the websocket client c.

Used for load monitoring.
*/
func (c WSClient) InBuffer() (buff comm.BuffStat) {
	return comm.BuffStat{
		Cap:  inBuffMax,
		Curr: c.inCount,
	}
}

/*
OutBuffer reports the status of the outgoing payload
buffer for the websocket client c.

Used for load monitoring.
*/
func (c WSClient) OutBuffer() (buff comm.BuffStat) {
	return comm.BuffStat{
		Cap:  outBuffMax,
		Curr: c.outCount,
	}
}

/*
Quit closes the websocket connection with the
peer. The client should be considered unregistered
from the multiplexer's collection of active
connections.
It always returns a non-nil error.
*/
func (c WSClient) Quit() (t comm.OPToken) {
	// propagate quit signal and close channel
	c.quit <- true
	close(c.quit)
	return t
}
