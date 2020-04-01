// Package connection defines and provides functionality
// for creating chat connections to the service for real
// time chat.
package chat

import (
	"encoding/json"
	"fmt"
	"log"

	ws "github.com/gorilla/websocket"

	"github.com/navaz-alani/go-talk/core/chat/payload"
)

// Connection defines an active connection to the
// chat service from a chat client. It manages data
// processing and transmission between the chat Service
// and the websocket connection.
type Connection struct {
	Service *Service
	Sock    *ws.Conn
	Owner   string
	ToSend  chan *payload.ChatItem
	Control chan *payload.ControlItem
}

// readFromSock is a routine which reads messages from
// the active websocket connection and passes them to the
// upstream service for distribution.
// All incoming messages are saved upon receipt.
func (c *Connection) readFromSock() {
	var unexpectedClose bool
	defer func() {
		if !unexpectedClose {
			c.Service.unregister <- c
			_ = c.Sock.Close()
		}
	}()

	for {
		msgType, msg, err := c.Sock.ReadMessage()
		if err != nil {
			log.Println("read err: ", err)
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway,
				ws.CloseAbnormalClosure) {
				c.Service.control <- payload.NewErr(c.Owner,
					"error: [sock] connection unexpectedly closed")
				unexpectedClose = true
			}
			break
		}

		if msgType != ws.TextMessage {
			err := payload.NewErr(c.Owner, "error: [sock] expected text payload")
			c.Service.control <- err
			continue
		}

		var pLoad struct {
			Type string `json:"type"`
			Data string `json:"data"`
		}
		if err := json.Unmarshal(msg, pLoad); err != nil {
			err := payload.NewErr(c.Owner, "error: decode fail")
			c.Service.control <- err
			continue
		}

		if pLoad.Type == payload.TypeControl {
			// create a control payload and send
			chatItem, err := payload.DecodeChatItem([]byte(pLoad.Data))
			if err != nil {
				c.Service.control <- payload.NewErr(c.Owner, err.Error())
			}
			c.Service.distribute <- chatItem
		} else if pLoad.Type == payload.TypeControl {
			// create a chat item payload and send
			ctrlItem, err := payload.DecodeControl([]byte(pLoad.Data))
			if err != nil {
				c.Service.control <- payload.NewErr(c.Owner, err.Error())
			}
			c.Service.control <- ctrlItem
		} else {
			err := payload.NewErr(c.Owner, "error: invalid payload type")
			c.Service.control <- err
		}
	}
}

// writeToSock is a routine which listens on the ToSend channel
// and forwards payloads to the chat client over the active
// websocket connection.
func (c *Connection) writeToSock() {
	for {
		select {
		case msg := <-c.ToSend:
			data, err := msg.Serialize()
			if err != nil {
				c.Service.control <- payload.NewErr(msg.From(),
					fmt.Sprintf("error: serialization fail, %s", err.Error()))
			}
			err = c.Sock.WriteMessage(ws.TextMessage, data)
			if err != nil {
				err := payload.NewErr(msg.From(),
					"error: [sock] failed to write")
				c.Service.control <- err
			}
		}
	}
}

// Listen is a routine which starts the read and write
// routines over the websocket connection for the client.
func (c *Connection) Listen() {
	go c.writeToSock()
	go c.readFromSock()
}
