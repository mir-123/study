package methods

import (
	"log"
)

// ————结构体官方案例开始————
// 假设有下面这样一个结构体
type Person struct {
	Name     string
	Age      int
	Address  string
	Salary   float64
	Birthday string
}

// 声明一个PersonOptions类型，它接受一个*Person类型的参数，它必须是指针，因为我们要在闭包中对Person赋值。
type PersonOptions func(p *Person) // todo 根据猜测 这种应该类似于 创建一个表 然后将表return

// 接下来创建选项函数，它们一般是With开头，它们的返回值就是一个闭包函数。
func WithName(name string) PersonOptions {
	return func(p *Person) {
		p.Name = name
	}
}

func WithAge(age int) PersonOptions {
	return func(p *Person) {
		p.Age = age
	}
}

func WithAddress(address string) PersonOptions {
	return func(p *Person) {
		p.Address = address
	}
}

func WithSalary(salary float64) PersonOptions {
	return func(p *Person) {
		p.Salary = salary
	}
}

// 实际声明的构造函数签名如下，它接受一个可变长PersonOptions类型的参数。
func NewPerson(options ...PersonOptions) *Person {
	// 优先应用options
	p := &Person{}
	for _, option := range options {
		option(p)
	}

	// 默认值处理
	if p.Age < 0 {
		p.Age = 0
	}

	return p
}

// ————结构体官方案例结束————

// ————map 映射表 官方案例开始————
func MapUser() {

	// map[string]后面可以跟各种类型

	// 3.键为字符串，值为另一个以字符串为键、整数为值的 map
	// 外层 map 的键是字符串，值是一个内层的 map
	cardMap3 := map[string]map[string]int{
		"aa": {
			"age":  10,
			"high": 120,
		},
		"bb": {
			"age":  30,
			"high": 180,
		},
	}
	log.Println(cardMap3, 55555555)               // map[aa:map[age:10 high:120] bb:map[age:30 high:180]] 55555555
	log.Println(cardMap3["aa"], 66666)            // map[age:10 high:120] 66666
	log.Println(cardMap3["aa"]["age"], 777777777) // 10 777777777

	// 2. 键是对象
	type Info struct {
		Name string
	}
	cardMap2 := map[string]Info{
		"123": {Name: "小猫"},
		"789": {Name: "小狗"},
		"345": {Name: "小兔"},
	}
	// 检查 "123" 对应的值
	value, exists := cardMap2["123"]
	log.Println(value, 77777)     // {小猫} 77777
	log.Println(exists, 88888888) // true 88888888

	// 1. 键是字符串 值也是字符串
	cardMap := map[string]string{
		"123": "小猫",
		"789": "小狗",
		"345": "小兔",
	}
	date, exists := cardMap["123"]
	log.Println(date, exists, 666) // 小猫 true 666

	date1, exists1 := cardMap["12333"]
	log.Println(date1, exists1, 777) // false 777
}

// ————map 映射表 官方案例结束————
