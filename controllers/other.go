package controllers

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
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

func (u UserController) GetEcodJson(c *gin.Context) {
	type Person struct {
		Name    string `json:"name"`          // 编码为 JSON 时使用 "name" 作为键
		Age     int    `json:"age,omitempty"` // 若 Age 为 0 则在 JSON 中省略
		Address string `json:"-"`             // 编码为 JSON 时忽略该字段
	}

	p1 := Person{Name: "Alice", Age: 25}
	p2 := Person{Name: "Bob", Age: 0, Address: "Somewhere"}

	jsonData1, _ := json.Marshal(p1)
	jsonData2, _ := json.Marshal(p2)

	fmt.Println(string(jsonData1)) // {"name":"Alice","age":25}
	fmt.Println(string(jsonData2)) // {"name":"Bob"}
}
