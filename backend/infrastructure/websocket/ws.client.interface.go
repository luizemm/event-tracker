package websocket

import "github.com/gorilla/websocket"

type WsClientInterface interface{
	SetConnection(conn *websocket.Conn)
	Execute()
}