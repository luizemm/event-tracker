package forward

import (
	"github.com/luizemm/data-collector/infrastructure/websocket/receiver"
	"github.com/luizemm/data-collector/usecase/event"
)

type forwardEventUseCase struct {
	senderManager *receiver.WsReceiverManager
}

func (f *forwardEventUseCase) Execute(eventDto event.EventDto) {
	f.senderManager.Send <- eventDto
}

func NewForwardEventUseCase(senderManager *receiver.WsReceiverManager) ForwardEventUseCaseInterface {
	return &forwardEventUseCase{
		senderManager: senderManager,
	}
}