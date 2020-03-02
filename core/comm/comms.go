/*
Package comm contains type definitions for objects
used for both internal (within the backend) and
external data transmission.
*/
package comm

/*
Payload is an interface describing a communications
payload for a message/interaction between.
It is used internally for object-invariant payload
processing.
*/
type Payload interface {
	/*
		Raw serializes the payload's contents to a
		slice of bytes which can be sent over a HTTP
		or websocket connection.
	*/
	Raw() []byte
}

/*
OPToken is an interface describing the response
to an internal operation.
*/
type OPToken interface {
	/*
		Error returns the error associated with
		the operation.
	*/
	Error() string
}

/*
BuffStat is used to report the statistics of
a buffered channel.
*/
type BuffStat struct {
	/*
		Cap is the buffer's maximum capacity.
	*/
	Cap int
	/*
		Curr is the buffered objects count.
	*/
	Curr int
}
