package sender

import (
	"github.com/luizemm/data-collector/infrastructure/log"
)

type WsSenderManager struct {
	clients map[*WsSenderClient]bool
	register chan *WsSenderClient
	unregister chan *WsSenderClient
}

func NewSenderManager() *WsSenderManager {
	return &WsSenderManager{
		clients: make(map[*WsSenderClient]bool),
		register: make(chan *WsSenderClient),
		unregister: make(chan *WsSenderClient),
	}
}

func (m *WsSenderManager) Run() {
	for {
		select {
		case client := <- m.register:
			m.clients[client] = true
			log.Logger.Debug("sender registered")
		case client := <- m.unregister:
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				log.Logger.Debug("sender unregistered")
			}
		}
	}
}