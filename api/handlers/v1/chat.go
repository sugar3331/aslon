package v1

import (
	"aslon/gee"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type ChatApi struct {
}

//func Chat(w http.ResponseWriter, r *http.Request) {
//	chatService.Chat(w, r)
//}

func P2PChat(w http.ResponseWriter, r *http.Request) {
	chatService.P2PChat(w, r)
}

func (ch *ChatApi) GetChat(c *gee.Context) {
	c.HTML(http.StatusOK, "chat.html", gee.H{})
}

func (ch *ChatApi) GetP2PChat(c *gee.Context) {
	c.HTML(http.StatusOK, "p2pChat.html", gee.H{})
}
