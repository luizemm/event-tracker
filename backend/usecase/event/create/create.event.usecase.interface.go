package create

import "github.com/luizemm/event-tracker/usecase/event"

type CreateEventUseCaseInterface interface {
	Execute(event event.EventDto)
}