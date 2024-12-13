package receiver

import (
	"github.com/luizemm/data-collector/infrastructure/log"
	"github.com/luizemm/data-collector/usecase/event"
)

type WsReceiverManager struct {
	clients map[*WsReceiverClient]bool
	register chan *WsReceiverClient
	unregister chan *WsReceiverClient

	Send chan event.EventDto
}

func NewReceiverManger() *WsReceiverManager{
	return &WsReceiverManager{
		clients: make(map[*WsReceiverClient]bool),
		register: make(chan *WsReceiverClient),
		unregister: make(chan *WsReceiverClient),
		Send: make(chan event.EventDto, 2),
	}
}

func (m *WsReceiverManager) Run() {
	for {
		select {
		case client := <- m.register:
			m.clients[client] = true
			log.Logger.Debug("receiver registered")
		case client := <- m.unregister:
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.send)
				log.Logger.Debug("receiver unregistered")
			}
		case event := <- m.Send:
			log.Logger.Debug("sending event to receiver clients")
			for client := range m.clients {
				client.send <- event
			}
		}
	}
}