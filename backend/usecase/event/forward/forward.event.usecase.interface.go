package forward

import "github.com/luizemm/data-collector/usecase/event"

type ForwardEventUseCaseInterface interface{
	Execute(event.EventDto)
}