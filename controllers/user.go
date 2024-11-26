package controllers

import "github.com/gin-gonic/gin"

type UserController struct{} // 避免同一个包下其他文件内的同名函数

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	ReturnSuccess(c, 0, "success", id, 1)
}

func (u UserController) GetList(c *gin.Context) {
	ReturnError(c, 403, "error")
}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (u UserController) PostUserInfo(c *gin.Context) {

	////form格式
	//cid := c.PostForm("cid")
	//name := c.DefaultPostForm("name", "wanghu") // 默认值
	//ReturnSuccess(c, 200, cid, name, 1)

	// json格式
	//param := make(map[string]interface{}) // todo interface{}不确定属性为什么可以是接口类型 make是啥
	//err := c.BindJSON(&param) // todo c.BindJSON(&param)
	//if err == nil {
	//	ReturnSuccess(c, 4000, param["name"], param[cid], 1)
	//	return
	//}
	//ReturnError(c, 4001, gin.H{"err": err}) // todo gin.H是啥

	// 结构体
	search := &Search{}
	err := c.BindJSON(&search)
	if err == nil {
		ReturnSuccess(c, 4000, search.Name, search.Cid, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err}) // todo gin.H是啥
}
