package controllers

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetCode 获取加密的密文 家医门诊大屏读取身份证后加密的一些操作
func (u UserController) GetCode(c *gin.Context) {
	// 生成密钥的测试
	swipe := struct {
		Card string
	}{
		Card: "410381199907264525",
	}

	now := time.Now().Unix()
	secret, err := EncryptAES([]byte(strconv.Itoa(int(now))+swipe.Card), []byte("16byte_key_examp"))
	if err != nil {
		log.Println(22222)
		return
	}
	code := hex.EncodeToString(secret)
	log.Println(code, 1111122)

	// 返回加密的密文
	c.JSON(http.StatusOK, struct {
		Code string
	}{code})
}
