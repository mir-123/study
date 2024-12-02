package main

import (
	"studyGIN/router"
)

func main() {
	r := router.Router()

	// 启动 Gin 服务器，监听 9999 端口
	r.Run(":8066")
}
