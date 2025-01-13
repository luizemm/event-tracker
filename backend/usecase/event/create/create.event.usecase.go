package create

import (
	"log/slog"

	"github.com/luizemm/event-tracker/domain/event"
	"github.com/luizemm/event-tracker/infrastructure/log"
	dto "github.com/luizemm/event-tracker/usecase/event"
)

type createEventUseCase struct {
	eventDb event.EventPersistenceInterface
}

func (u *createEventUseCase) Execute(eventDto dto.EventDto) {
	event := event.NewEvent(event.EventProps{
		EventType: eventDto.EventType,
		Data: eventDto.Data,
		Timestamp: eventDto.Timestamp,
	})

	err := u.eventDb.Save(event)

	if err != nil {
		log.Logger.Error("Save event in database", slog.Any("error", err))
	}
}

func NewCreateEventUseCase(eventDb event.EventPersistenceInterface) CreateEventUseCaseInterface {
	return &createEventUseCase{
		eventDb: eventDb,
	}
}