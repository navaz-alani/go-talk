// Package payload defines and provides functionality
// related to data transmission over an active websocket
// connection with a chat client.
package payload

// data types
const (
	TypeControl = "type-control"
	TypeText    = "type-text"
)

// Payload is a type which can be used for data
// transmission over a websocket connection.
type Payload interface {
	// Type specifies the type of the payload.
	Type() string
	// Dest specifies the destination of the data.
	Dest() string
	// Data is the content of the payload which
	// has to be delivered.
	Data() interface{}
	// Kind specified the kind of payload
	Kind() string
}
