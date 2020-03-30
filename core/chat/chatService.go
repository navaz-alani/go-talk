package chat

import (
	"github.com/navaz-alani/go-talk/core/chat/payload"
)

type UID = string

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
		case ci := <- s.control:
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
	chatService := Service{
		Active:     make(map[UID]*Connection),
		register:   make(chan *Connection),
		unregister: make(chan *Connection),
		distribute: make(chan *payload.ChatItem),
		control:    make(chan *payload.ControlItem),
	}

	go chatService.Run()
}
