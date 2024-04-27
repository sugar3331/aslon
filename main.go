package main

import (
	v1 "aslon/api/handlers/v1"
	"aslon/core"
	"aslon/global"
	"aslon/initialize"
	"aslon/service"
	"fmt"
	"log"
	"net/http"
)

func main() {

	global.Viper = core.Viper()               //初始化viper
	global.UserDB = initialize.GormTalMysql() //初始化mysql连接
	global.Mogo = initialize.MongodbInit()    //初始化mongodb
	global.SnowflakeNode = core.SnowFlake()   //初始化雪花算法节点

	go func() {
		http.HandleFunc("/publicChat", service.PublicChat)
		http.HandleFunc("/chat", service.Chat)
		http.HandleFunc("/p2pChat", v1.P2PChat)
		log.Printf("IM Websocket Starting server at: %v...\n", global.GlobalConfig.System.ChatAddr)
		chatAddr := fmt.Sprintf(":%v", global.GlobalConfig.System.ChatAddr)
		err := http.ListenAndServe(chatAddr, nil)
		if err != nil {
			log.Fatal("聊天服务启动失败！！！:", err)
		}
	}()

	core.RunWindowsServer()
}
