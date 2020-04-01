// Package chat provides the chat functionality for
// the core service of the application.
package chat

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/navaz-alani/go-talk/core/auth"
)

const IDLen = 25

const (
	TypePrivate = "private"
)

// Chat defines the structure of a chat.
type Chat struct {
	ID           primitive.ObjectID `json:"-" bson:"_id" _id_:"chat"`
	ChatID       string             `json:"chatID" _ax_:"true" _ix_:"true"`
	Type         string             `json:"type"`
	Participants []UID              `json:"participants"`
}

// NewFromControl takes a control message to create
// a new chat and creates a new chat from the data
// held within.
func NewFromControl() *Chat {
	return &Chat{
		ChatID:       auth.Key(IDLen),
		Type:         "",
		Participants: nil,
	}
}
