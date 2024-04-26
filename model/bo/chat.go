package bo

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	Conn     *websocket.Conn
	Username string
	Userid   string
}

type FullMessage struct {
	Message  string
	Username string
	Userid   string
	SendTime time.Time
}
