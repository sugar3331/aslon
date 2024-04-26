package service

import (
	"aslon/global"
	"aslon/middleware"
	"aslon/model/bo"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"sync"
	"time"
)

type ChatService struct {
}

var clients = &sync.Map{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Chat(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	userid, username, err := middleware.Token.ImJwtAuthMiddleware(token)
	if err != nil {
		return
	}
	//升级http连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建 Client 结构体，并将连接和用户名保存其中
	client := &bo.Client{
		Conn:     conn,
		Userid:   userid,
		Username: username,
	}

	// 将连接添加到clients列表
	clients.Store(client, true)
	defer func() {
		conn.Close()
		// 连接断开时从clients列表移除该连接
		clients.Delete(client)
	}()
	//从mongodb里面读缓存
	go sendRecentMessages(client)
	// 处理WebSocket连接
	for {
		// 读取消息
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("Received message:", string(msg))
		sendtimeNow := time.Now()
		fullMe := bo.FullMessage{
			Message:  string(msg),
			Username: username,
			Userid:   userid,
			SendTime: sendtimeNow,
		}
		log.Println("开始发消息，用户id: " + userid + ":" + string(msg))
		go saveMessageToMongoDB(fullMe)
		// 广播消息给所有客户端
		go broadcastMessage(messageType, fullMe, client)
	}
}

func sendRecentMessages(client *bo.Client) {
	opts := options.Find().SetSort(bson.D{{"_id", -1}}).SetLimit(30)
	filter := bson.D{{}}
	collection := global.Mogo.Database("ImChat").Collection("publicChat")
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err)
		return
	}

	var recentMessages []bo.FullMessage
	for cursor.Next(context.TODO()) {
		var message bo.FullMessage
		if err := cursor.Decode(&message); err != nil {
			log.Println(err)
			break
		}
		recentMessages = append(recentMessages, message)
	}

	// 倒序发送最近的 30 条记录给客户端
	for i := len(recentMessages) - 1; i >= 0; i-- {
		err := client.Conn.WriteJSON(recentMessages[i])
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func saveMessageToMongoDB(msg bo.FullMessage) error {
	collection := global.Mogo.Database("ImChat").Collection("publicChat")
	_, err := collection.InsertOne(context.Background(), msg)
	if err != nil {
		return err
	}
	return nil
}

// broadcastMessage 服务端把用户发送的消息推送给所有在线用户的广播函数
func broadcastMessage(messageType int, fullMe bo.FullMessage, sender *bo.Client) {
	clients.Range(func(key, value interface{}) bool {
		client := key.(*bo.Client)

		// 不向消息发送者推送消息
		if client == sender {
			return true
		}

		err := client.Conn.WriteJSON(fullMe)
		if err != nil {
			log.Println(err)
			client.Conn.Close()
			clients.Delete(client)
		}
		return true
	})
}

var p2pclients sync.Map

func (ch *ChatService) P2PChat(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	log.Println("开始进行点对点通信，首先身份验证")
	userid, username, err := middleware.Token.ImJwtAuthMiddleware(token)
	if err != nil {
		return
	}

	targetUserID := parseTargetUserIDFromURL(r)
	if userid == targetUserID {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"error": "不能和自己聊天"}
		json.NewEncoder(w).Encode(response)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade websocket:", err)
		return
	}

	//username := r.Context().Value("userName").(string)
	p2pclients.Store(userid, conn)
	// 连接关闭时，从映射中删除用户ID和连接
	exitSignal := make(chan struct{})
	defer func() {
		conn.Close()
		p2pclients.Delete(userid) // 使用 Delete 方法来删除键值对，保证并发安全性
		close(exitSignal)
	}()

	log.Println("身份验证完毕用户id: " + userid + " 用户: " + username + " 开始与对方id通信: " + targetUserID)
	go sendP2PRecentMessages(userid, targetUserID, conn)
	go handleIncomingMessage(userid, targetUserID, username, conn, exitSignal)

	// 在这里等待信号通知退出
	<-exitSignal
}

func handleIncomingMessage(userid, targetUserID, username string, conn *websocket.Conn, exitSignal <-chan struct{}) {
	for {
		select {
		case <-exitSignal:
			return
		default:
			_, rep2pmessage, err := conn.ReadMessage()
			if err != nil {
				log.Println("Failed to read message from websocket:", err)
				return
			}
			log.Println("Received message:", string(rep2pmessage))
			sendtimeNow := time.Now()
			pup2pmessage := bo.FullMessage{
				Message:  string(rep2pmessage),
				Username: username,
				Userid:   userid,
				SendTime: sendtimeNow,
			}
			go func() {
				err := saveP2PMessageToMongoDB(userid, targetUserID, pup2pmessage)
				if err != nil {
					log.Println("mongodbp2p 存储消息有误 ", err)
				}
			}()
			// 转发消息给目标用户
			if targetConn, ok := p2pclients.Load(targetUserID); ok {
				err = targetConn.(*websocket.Conn).WriteJSON(pup2pmessage)
				if err != nil {
					log.Println("Failed to send message to target user:", err)
				}
			} else {
				log.Println("Target user is not connected.")
			}
		}
	}
}

func parseTargetUserIDFromURL(r *http.Request) string {
	return r.URL.Query().Get("taruserid")
}

func sendP2PRecentMessages(userid, targetUserID string, conn *websocket.Conn) {
	opts := options.Find().SetSort(bson.D{{"_id", -1}}).SetLimit(30)
	filter := bson.D{{}}
	collectionName := convert(userid, targetUserID)
	collection := global.Mogo.Database("ImP2PChat").Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err)
		return
	}

	var recentMessages []bo.FullMessage
	for cursor.Next(context.TODO()) {
		var message bo.FullMessage
		if err := cursor.Decode(&message); err != nil {
			log.Println(err)
			break
		}
		recentMessages = append(recentMessages, message)
	}

	// 倒序发送最近的 30 条记录给客户端
	for i := len(recentMessages) - 1; i >= 0; i-- {
		//if targetConn, ok := p2pclients.Load(userid); ok {
		//	err = targetConn.(*websocket.Conn).WriteJSON(recentMessages[i])
		err = conn.WriteJSON(recentMessages[i])
		if err != nil {
			log.Println(err)
			break
		}

	}
}

func saveP2PMessageToMongoDB(userid, targetUserID string, msg bo.FullMessage) error {
	collectionName := convert(userid, targetUserID)
	collection := global.Mogo.Database("ImP2PChat").Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), msg)
	if err != nil {
		return err
	}
	return nil
}

func convert(a, b string) string {
	if a > b {
		return a + "-" + b
	} else {
		return b + "-" + a
	}
}
