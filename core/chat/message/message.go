// Package message defines a Message type which
// represents a single contribution to a chat from
// a specific client.
package message

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/navaz-alani/go-talk/core/auth"
	"github.com/navaz-alani/go-talk/core/chat/payload"
)

const IDLen = 25

const (
	Text = "text"
)

type MsgID = string

// Message defines the structure of a message in
// the application's core service.
type Message struct {
	ID    primitive.ObjectID `json:"-" bson:"_id" _id_:"message"`
	MsgID MsgID              `json:"msgID" _ax_:"true" _ix_:"true"`
	Kind  string             `json:"type"`
}

// NewFromPayload creates a message from the given
// message payload. The given message is completely
// initialized and saved in persistent storage.
func NewFromPayload(p *payload.ChatItem) *Message {
	return &Message{
		MsgID: auth.Key(IDLen),
		Kind:  p.Kind(),
	}
}
