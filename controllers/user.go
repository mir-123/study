package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct{} // 避免同一个包下其他文件内的同名函数

// GetTest 仅用于测试
func (u UserController) GetTest(c *gin.Context) {
	numPtr := new(int)
	fmt.Println(numPtr)

	//row := struct {
	//	Name    string
	//	NewUser bool
	//}{"你好", true}
	//
	//c.JSON(http.StatusOK, row)
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
