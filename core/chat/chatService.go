package chat

import (
	"log"
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
			log.Println("<- registered")
			s.Active[c.Owner] = c
		case c := <-s.unregister:
			log.Println("<- unregistered")
			delete(s.Active, c.Owner)
		case p := <-s.distribute:
			dest := (*p).Dest()
			if destConn := s.Active[dest]; destConn != nil {
				destConn.ToSend <- p
			}
		case ci := <-s.control:
			if destConn := s.Active[ci.Dest()]; destConn != nil {
				destConn.Control <- ci
			}
		}
	}
}

// Init initializes the core component's chat
// service.
func Init() {
	ChatService = &Service{
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
	log.Println("connection request received.")
	c, err := (&ws.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
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

	conn.Listen()
	ChatService.register <- conn
}
