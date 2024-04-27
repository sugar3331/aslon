package initialize

import (
	"aslon/global"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func MongodbInit() *mongo.Client {
	// 设置凭据
	credential := options.Credential{
		Username: global.GlobalConfig.Mongo.Username,
		Password: global.GlobalConfig.Mongo.Password,
	}
	url := fmt.Sprintf("mongodb://%s:%s", global.GlobalConfig.Mongo.Path, global.GlobalConfig.Mongo.Port)
	// 设置连接选项
	clientOptions := options.Client().ApplyURI(url).SetAuth(credential)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// 处理连接失败的错误
		log.Println("mongodb 连接数据库失败:", err)
	} else {
		log.Println("mongodb 连接数据库成功")
	}
	return client
}
