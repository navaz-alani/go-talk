package chat

import (
	"net/http"

	ws "github.com/gorilla/websocket"

	"github.com/navaz-alani/go-talk/core/chat/payload"
)

type UID = string

var ChatService *Service

type Service struct {
	Active map[UID]*Connection

	// Service channels |
	//                  V

	register   chan *Connection
	unregister chan *Connection
	distribute chan *payload.ChatItem
	control    chan *payload.ControlItem
}

func (s Service) Run() {
	for {
		select {
		case c := <-s.register:
			s.Active[c.Owner] = c
		case c := <-s.unregister:
			delete(s.Active, c.Owner)
		case p := <-s.distribute:
			dest := (*p).Dest()
			if destConn := s.Active[dest]; destConn != nil {
				destConn.ToSend <- p
			}
		case ci := <-s.control:
			switch ci.Kind() {
			case payload.CtrlNewChat:
				// create a new chat

			}
		}
	}
}

// Init initializes the core component's chat
// service.
func Init() {
	ChatService := &Service{
		Active:     make(map[UID]*Connection),
		register:   make(chan *Connection),
		unregister: make(chan *Connection),
		distribute: make(chan *payload.ChatItem),
		control:    make(chan *payload.ControlItem),
	}

	go ChatService.Run()
}

// NewConnection takes a request to initiate a connection with
// the chat service and upgrades it to a websocket connection
// which can be for real-time chat.
func NewConnection(w http.ResponseWriter, r *http.Request) {
	c, err := (&ws.Upgrader{}).Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "error: [chat] failed to upgrade protocol",
			http.StatusInternalServerError)
	}

	uid := r.Context().Value("uid").(string)
	conn := &Connection{
		Service: ChatService,
		Sock:    c,
		Owner:   uid,
		ToSend:  make(chan *payload.ChatItem),
		Control: make(chan *payload.ControlItem),
	}

	ChatService.register <- conn
}
