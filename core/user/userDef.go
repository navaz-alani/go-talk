// Package user supplies functionality for
// users in the context of the application.
package user

import (
	"github.com/navaz-alani/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users is a handle to the persistent storage
// for the application's users.
var Users *entity.Entity

// User defines the structure of user in the
// context of this application.
type User struct {
	ID       primitive.ObjectID `json:"-" bson:"_id" _id_:"user"`
	UID      string             `json:"-" bson:"uid" _ax_:"true"`
	Name     string             `json:"name" _hd_:"c"`
	Username string             `json:"username" _hd_:"c" _ax_:"true"`
	Email    string             `json:"email" _hd_:"c" _ax_:"true"`
	PwdHash  string             `json:"-" bson:"pwdHash"`
}
