package core

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func SnowFlake() (Node *snowflake.Node) {
	// 初始化 Snowflake 节点
	var err error
	Node, err = snowflake.NewNode(1)
	if err != nil {
		// 处理错误
		fmt.Println("Error occurred:", err)
	}
	return Node
}
