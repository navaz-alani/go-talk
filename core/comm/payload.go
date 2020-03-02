package comm

type WSPayloadType int

const (
	WSMsg WSPayloadType = 1
	WSBin WSPayloadType = 2
)

/*
WSPayload is a type used for transmission of
information between web-socket connections.
*/
type WSPayload struct {
	/*
		Signal is a string which specifies the
		type of the payload (e.g. "quit").
	*/
	Signal string
	/*
		Type specifies the type of the payload
		contained within.
	*/
	Type WSPayloadType
	Str  string
	Bin  []byte
}

/*
Raw returns a slice of bytes encoding the
payload data.
*/
func (m WSPayload) Raw() []byte {
	if m.Type == WSMsg {
		return []byte(m.Str)
	}
	return m.Bin
}
