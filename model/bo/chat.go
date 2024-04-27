package bo

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	Conn     *websocket.Conn
	UserName string
	UserId   string
	UserNick string
}

type FullMessage struct {
	Message  string
	UserName string
	UserId   string
	UserNick string
	SendTime time.Time
}
