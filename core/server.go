package core

import (
	"aslon/global"
	"aslon/initialize"
	"fmt"
	"log"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GlobalConfig.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	log.Printf("server run success on %s\n", address)

	fmt.Printf(`
	----------------------------------------------------
	----------------------------------------------------
	----------------------------------------------------
	项目内部启动端口：%s 
	mysql数据库：地址：%s ,端口号：%s
	----------------------------------------------------
	----------------------------------------------------
	----------------------------------------------------
`, address,
		global.GlobalConfig.MysqlUser.Path, global.GlobalConfig.MysqlUser.Port,
	)

	fmt.Println(s.ListenAndServe().Error())
}
