package controllers

import "github.com/gin-gonic/gin"

type JsonStruct struct {
	Code  int         `json:"code"`  // 状态码
	Msg   interface{} `json:"msg"`   // 可以输字符串或者数字
	Data  interface{} `json:"data"`  // 列表 不确定类型 所用用 interface
	Count int64       `json:"count"` // 总数 直接从数据库取出的是 int64所以此处也用int64
}

type JsonErrStruct struct {
	Code int         `json:"code"` // 状态码
	Msg  interface{} `json:"msg"`  // 可以输字符串或者数字
}

// ReturnSuccess 正确的时候
func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(200, json)
}

// ReturnError 错误的时候
func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(400, json)
}
