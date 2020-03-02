package core

import (
	"net/http"

	muxCtx "github.com/navaz-alani/entity/multiplexer/muxContext"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"

	"github.com/navaz-alani/go-talk/core/auth"
	"github.com/navaz-alani/go-talk/core/client"
	"github.com/navaz-alani/go-talk/core/comm"
)

/*
User defines the structure of a user of the go-talk
service.
*/
type User struct {
	/*
		ID is the user's database ID.
	*/
	ID primitive.ObjectID `json:"-" bson:"_id" _id_:"user"`
	/*
		Name is the user's full name.
	*/
	Name string `json:"name" bson:"name" _hd_:"c"`
	/*
		Username is the unique (user-chosen) identifier
		for the user.
	*/
	Username string `json:"username" bson:"u-name" _ax_:"true" _ix_:"true" _hd_:"c"`
	/*
		Dist is a slice of comm.Payloads which have to
		be distributed to the user upon connection.
	*/
	Dist []comm.Payload `json:"-" bson:"dist"`
}

/*
RegisterUser handles requests to create a new user for
the go-talk service.
It accepts a JSON payload defining the fields tagged with
the multiplexer.CreationFieldsToken in the eField.HandleTag.
All other information provided is ignored.
*/
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// isolate user from context
	muxContext, _ := muxCtx.IsolateCtx(r)
	usr := muxContext.Retrieve("user").(User)

	// add to database
	addID, err := EMux.E("user").Add(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create JWT with ID
	tok, err := auth.EncodePayload(addID.Hex())
	if err != nil {
		http.Error(w, "JWT gen fail: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// respond with new session token (JWT)
	_, err = w.Write([]byte(tok))
	if err != nil {
		http.Error(w, "response write fail", http.StatusInternalServerError)
		return
	}
}

/*
Connect handles requests to establish a new client connection
with the core service. It upgrades the connection to a
web-socket connection and internally registers the client.
*/
func Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := client.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "protocol upgrade fail: "+err.Error(), http.StatusBadRequest)
		return
	}

	wsc := client.NewWSClient(conn)
	c := client.Client(wsc)
	CMux.RegisterClient(&c)

	// spawn listener routine
	go func() {
		incoming, quit := c.Listen()

		for {
			select {
			case payload := <-incoming:
				c.Send(context.TODO(), payload)
			case <-quit:
				conn.Close()
			}
		}
	}()
}
