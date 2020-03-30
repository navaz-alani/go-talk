package payload

import (
	"encoding/json"
	"fmt"
)

// ControlItem data kinds
const (
	CtrlNewChat = "new-chat"
	CtrlErr     = "err"
)

// ControlItem is a payload implementation which is used to
// transmit status updates and chat commands over an active
// websocket connection to a chat client.
// ControlItems are NOT saved to persistent storage.
//
// ControlItem payloads are generated and used by the app,
// in response to actions by the chat client.
type ControlItem struct {
	kind string
	dest string
	data string
}

func (cp *ControlItem) Type() string      { return TypeControl }
func (cp *ControlItem) Dest() string      { return cp.dest }
func (cp *ControlItem) Data() interface{} { return cp.data }
func (cp *ControlItem) Kind() string      { return cp.kind }
func (cp *ControlItem) IsErr() bool       { return cp.kind == CtrlErr }

// DecodeControl decodes a control message from the given
// data. If decoding is successful, expect a nil error.
func DecodeControl(data []byte) (*ControlItem, error) {
	var decoder struct {
		Kind string `json:"kind"`
		Dest string `json:"dest"`
		Data string `json:"data"`
	}

	if err := json.Unmarshal(data, &decoder); err != nil {
		return nil, err
	}

	return &ControlItem{
		kind: decoder.Kind,
		dest: decoder.Dest,
		data: decoder.Data,
	}, nil
}

// NewControl creates and returns a pointer to a
// control payload.
// kind is a finer classification of the ControlItem payload.
// dest is the UID of the user to which this control payload
// is to be delivered.
// Finally, data is a string containing a control message or
// a control command.
func NewControl(kind, dest, data string) *ControlItem {
	return &ControlItem{
		kind: kind,
		dest: dest,
		data: data,
	}
}

// NewErr creates a control payload which corresponds to
// an error.
func NewErr(dest, data string) *ControlItem {
	return NewControl(CtrlErr, dest, data)
}

// Err returns the error associated with the control
// payload (returns nil if none).
func (cp *ControlItem) Err() error {
	if cp.IsErr() {
		return fmt.Errorf(cp.data)
	}
	return nil
}

// NewChat creates a control item to create a new chat.
func NewChat(data string) *ControlItem {
	return &ControlItem{
		kind: CtrlNewChat,
		data: data,
	}
}
