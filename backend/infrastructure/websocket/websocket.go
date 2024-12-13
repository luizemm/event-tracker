package websocket

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"regexp"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"github.com/luizemm/data-collector/infrastructure/db"
	"github.com/luizemm/data-collector/infrastructure/env"
	"github.com/luizemm/data-collector/infrastructure/log"
	"github.com/luizemm/data-collector/infrastructure/websocket/receiver"
	"github.com/luizemm/data-collector/infrastructure/websocket/sender"
	"github.com/luizemm/data-collector/usecase/event"
	"github.com/luizemm/data-collector/usecase/event/create"
	"github.com/luizemm/data-collector/usecase/event/forward"
)

func getPort() string {
    port := os.Getenv(env.PORT)

    if port == "" {
        port = "8080"
    }

    matched, err := regexp.MatchString("\\d+", port)

    if !matched {
        log.Logger.Error("Port must be a number")
        os.Exit(1)
    }

    if err != nil {
        log.Logger.Error("Regex port", slog.Any("error", err))
        os.Exit(1)
    }

    log.Logger.Info("port: " + port)

    return port
}

func Init(database *sql.DB) {	
	eventDb := db.NewEventDb(database)
	
	senderManager := sender.NewSenderManager()
	receiverManager := receiver.NewReceiverManger()

	forwardEventUseCase := forward.NewForwardEventUseCase(receiverManager)
	createEventUseCase := create.NewCreateEventUseCase(eventDb)

	go senderManager.Run()
	go receiverManager.Run()

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {return true},
	}

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		typ := r.URL.Query().Get("type")
        
        if typ == "" {
            w.WriteHeader(400)
            w.Write([]byte("Query parameter \"type\" is required"))
            return
        }

		var wsClient WsClientInterface

        switch typ {
        case "receiver":
            wsClient = receiver.NewReceiverClient(receiver.WsReceiverClientProps{
				SenderManager: receiverManager,
				Send: make(chan event.EventDto, 200),
			})
		case "sender":
			wsClient = sender.NewSenderClient(sender.WsSenderClientProps{
				ReceiverManager: senderManager,
				ForwardEventUseCase: forwardEventUseCase,
				CreateEventUseCase: createEventUseCase,
			})
		default:
			w.WriteHeader(400)
            w.Write([]byte("Unsupported type: " + typ))
			return
        }

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Logger.Error("Websocket upgrader", slog.Any("error", err))
			return
		}

		wsClient.SetConnection(conn)
		wsClient.Execute()
	})

    err := http.ListenAndServe(":" + getPort(), nil)

	if err != nil {
		log.Logger.Error("Listen server", slog.Any("error", err))
        os.Exit(1)
	}
}