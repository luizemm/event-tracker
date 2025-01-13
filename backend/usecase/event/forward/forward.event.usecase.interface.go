package forward

import "github.com/luizemm/event-tracker/usecase/event"

type ForwardEventUseCaseInterface interface{
	Execute(event.EventDto)
}