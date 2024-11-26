package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
	"studyGIN/repositories/methods"
)

// 根据json文件 导出表格
func (u UserController) ExportTable(c *gin.Context) {
	// 读取文件
	basePath := "D:\\Mir\\test\\aaa\\export\\in_1125\\"
	for i := 67; i <= 314; i++ {
		filePath := fmt.Sprintf("%s%d.txt", basePath, i)
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("打开文件出错:", err)
			return
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("读取文件内容出错:", err)
			return
		}
		// 处理读取到的文件
		//methods.JsonToExcel(string(data), strconv.Itoa(i))
		methods.JsonToExcel(string(data), "1125", strconv.Itoa(i))
	}
}
