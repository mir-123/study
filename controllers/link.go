package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// 获取连接时间
func (u UserController) GetLinkTime(c *gin.Context) {
	rsp := gin.H{
		"access_time": time.Now().Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, rsp)
}

// 获取链接发送过来的内容并返回
func (u UserController) GetLinkContent(c *gin.Context) {
	input := make(map[string]string)
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		if len(values) > 0 {
			input[key] = values[0]
		}
	}
	c.JSON(http.StatusOK, input)
}

func (u UserController) GetLinkWS(c *gin.Context) {
	// 升级到 WebSocket 连接
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("升级到 WebSocket 失败:", err)
		return
	}
	defer conn.Close()

	for {
		// 读取客户端发送的消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息错误:", err)
			return
		}

		// 处理消息，这里只是简单地返回收到的消息
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("发送消息错误:", err)
			return
		}

		time.Sleep(5 * time.Second) // 模拟处理耗时
	}
}
