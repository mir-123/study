package controllers

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"studyGIN/repositories/methods"
	"time"
)

type UserController struct{} // 避免同一个包下其他文件内的同名函数

// GetTest 仅用于测试
func (u UserController) GetTest(c *gin.Context) {
	methods.ArrCut()
}

// 密钥测试
func getSwipeKey() {
	// 生成密钥的测试
	swipe := struct {
		Card string
	}{
		Card: "410311111111111111",
	}

	now := time.Now().Unix()
	secret, err := EncryptAES([]byte(strconv.Itoa(int(now))+swipe.Card), []byte("16byte_key_examp"))
	if err != nil {
		log.Println(22222)
		return
	}
	code := hex.EncodeToString(secret)
	log.Println(code, 1111122)
}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	ReturnSuccess(c, 0, "success", id, 1)
}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (u UserController) PostUserInfo(c *gin.Context) {
	// 结构体
	search := &Search{}
	err := c.BindJSON(&search)
	if err == nil {
		ReturnSuccess(c, 4000, search.Name, search.Cid, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err}) // todo gin.H是啥
}
