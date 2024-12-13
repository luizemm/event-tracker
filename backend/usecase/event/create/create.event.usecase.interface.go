package create

import "github.com/luizemm/data-collector/usecase/event"

type CreateEventUseCaseInterface interface {
	Execute(event event.EventDto)
}