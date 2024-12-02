package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{} // 避免同一个包下其他文件内的同名函数

// GetTest 仅用于测试
func (u UserController) GetTest(c *gin.Context) {

	//// 显式组合的方式
	//type Person struct {
	//	name string
	//	age  int
	//}
	//
	//type Student struct {
	//	//p      Person // 显试
	//	Person // 隐式
	//	school string
	//}
	//
	//type Employee struct {
	//	p   Person
	//	job string
	//}
	//
	//// 在使用时需要显式的指定字段p
	////student := Student{
	////	p:      Person{name: "jack", age: 18},
	////	school: "lili school",
	////}
	////fmt.Println(student) // 打印结果：{{jack 18} lili school}
	//
	//// 匿名字段的名称默认为类型名，调用者可以直接访问该类型的字段和方法，但除了更加方便以外与第一种方式没有任何的区别。
	//student2 := Student{
	//	Person: Person{name: "jack", age: 18},
	//	school: "lili school",
	//}
	//fmt.Println(student2.Person.name) // jack
	//fmt.Println(student2.name)        // jack

	//// 测试 结构体 官方案例 用法
	//p1 := methods.NewPerson(
	//	methods.WithName("John Doe"),
	//	methods.WithAge(25),
	//	methods.WithAddress("123 Main St"),
	//	methods.WithSalary(10000.00),
	//)
	//p2 := methods.NewPerson(
	//	methods.WithName("Mike jane"),
	//	methods.WithAge(30),
	//)
	//p3 := methods.NewPerson()
	//log.Println(p1, 1111) // &{John Doe 25 123 Main St 10000 } 1111
	//log.Println(p2, 2222) // &{Mike jane 30  0 } 2222
	//log.Println(p3, 3333) // &{ 0  0 } 3333

	//numPtr := new(int)
	//fmt.Println(numPtr)

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
