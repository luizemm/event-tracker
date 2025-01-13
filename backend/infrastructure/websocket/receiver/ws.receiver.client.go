package receiver

import (
	encJson "encoding/json"

	"github.com/gorilla/websocket"
	"github.com/luizemm/event-tracker/usecase/event"
)

type WsReceiverClient struct {
    conn *websocket.Conn
    receiverManager *WsReceiverManager
    send chan event.EventDto
}

type WsReceiverClientProps struct {
	Conn *websocket.Conn
    ReceiverManager *WsReceiverManager
	Send chan event.EventDto
}

func NewReceiverClient(props WsReceiverClientProps) *WsReceiverClient {
	return &WsReceiverClient {
		conn: props.Conn,
		receiverManager: props.ReceiverManager,
		send: props.Send,
	}
}

func (c *WsReceiverClient) SetConnection(conn *websocket.Conn) {
	c.conn = conn
}

func (c *WsReceiverClient) Execute() {
	c.receiverManager.register <- c
	defer func() {
        c.receiverManager.unregister <- c
		c.conn.Close()
	}()
	for {
		select {
		case event, ok := <- c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte("Closing websocket"))
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

            json, err := encJson.Marshal(event)

            if err != nil {
				return
			}

			w.Write(json)

			n := len(c.send)
			for i := 0; i < n; i++ {
                json, err = encJson.Marshal(<-c.send)

                if err != nil {
                    return
                }

				w.Write([]byte("\n"))
				w.Write(json)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}