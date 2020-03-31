package payload

import (
	"encoding/json"
)

// ChatItem is a payload implementation used to
// transmit chat data over an active websocket
// connection to a chat client.
// Such a payload is intended to be converted to
// a message, saved to persistent storage and finally
// linked to a conversation.
type ChatItem struct {
	kind string
	dest string
	from string
	data interface{}
}

func (ci *ChatItem) Type() string      { return TypeText }
func (ci *ChatItem) Dest() string      { return ci.dest }
func (ci *ChatItem) Data() interface{} { return ci.data }
func (ci *ChatItem) Kind() string      { return ci.kind }
func (ci *ChatItem) From() string      { return ci.from }

func (ci *ChatItem) Serialize() ([]byte, error) {
	if data, err := json.Marshal(*ci); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

// NewChatItem creates a new ChatItem payload.
// kind is a finer classification of the ChatItem payload.
// dest is the ChatID of the chat under which the payload
// should be linked.
// data is the content of the payload.
func NewChatItem(kind, dest string, data interface{}) *ChatItem {
	return &ChatItem{
		kind: kind,
		dest: dest,
		data: data,
	}
}

// DecodeChatItem attempts to decode a ChatItem from the
// given data. If successful, the error is nil.
func DecodeChatItem(data []byte) (*ChatItem, error) {
	var decoder struct {
		Kind string      `json:"kind"`
		Dest string      `json:"dest"`
		From string      `json:"from"`
		Data interface{} `json:"data"`
	}

	if err := json.Unmarshal(data, &decoder); err != nil {
		return nil, err
	}

	return &ChatItem{
		kind: decoder.Kind,
		dest: decoder.Dest,
		from: decoder.From,
		data: decoder.Data,
	}, nil
}
