package sender

import (
	"encoding/json"
	"log/slog"
	"time"

	"github.com/gorilla/websocket"
	"github.com/luizemm/data-collector/infrastructure/log"
	"github.com/luizemm/data-collector/usecase/event"
	uCreate "github.com/luizemm/data-collector/usecase/event/create"
	uForward "github.com/luizemm/data-collector/usecase/event/forward"
)

type WsSenderClient struct {
    conn *websocket.Conn

    receiverManager *WsSenderManager

	createEventUseCase uCreate.CreateEventUseCaseInterface
	forwardEventUseCase uForward.ForwardEventUseCaseInterface
}

type WsSenderClientProps struct {
	Conn *websocket.Conn
    ReceiverManager *WsSenderManager
	CreateEventUseCase uCreate.CreateEventUseCaseInterface
	ForwardEventUseCase uForward.ForwardEventUseCaseInterface
}

func NewSenderClient(props WsSenderClientProps) *WsSenderClient {
	return &WsSenderClient{
		conn: props.Conn,
		receiverManager: props.ReceiverManager,
		createEventUseCase: props.CreateEventUseCase,
		forwardEventUseCase: props.ForwardEventUseCase,
	}
}

func (c *WsSenderClient) SetConnection(conn *websocket.Conn) {
	c.conn = conn
}

func (c *WsSenderClient) Execute() {
	c.receiverManager.register <- c
    defer func() {
		c.receiverManager.unregister <- c
		c.conn.Close()
	}()
	for {
		var eventDto event.EventDto

		_, message, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err) {
				log.Logger.Error("Receiver client closed", slog.Any("error", err))
				break
			}
			
			log.Logger.Error("Event input", slog.Any("error", err))
			continue
		}

		err = json.Unmarshal(message, &eventDto)

		if err != nil {
			log.Logger.Error("Invalid JSON format", slog.Any("error", err))
			
			errorMap := make(map[string]string)
			errorMap["timestamp"] = time.Now().Format(time.UnixDate)
			errorMap["error"] = "Invalid JSON format: " + err.Error()
			c.conn.WriteJSON(errorMap)
			continue
		}

		err = eventDto.Validate()

		if err != nil {
			errorMap := make(map[string]string)
			errorMap["timestamp"] = time.Now().Format(time.UnixDate)
			errorMap["error"] = err.Error()
			c.conn.WriteJSON(errorMap)
			continue
		}

		log.Logger.Info("Event: " + eventDto.Json())
		
		go c.createEventUseCase.Execute(eventDto)
		go c.forwardEventUseCase.Execute(eventDto)
	}
}