package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"io"
	"log"
	"os"
	"strconv"
	"studyGIN/repositories/methods"
)

// 根据json文件 导出表格
func (u UserController) ExportTable(c *gin.Context) {
	// 读取文件
	//basePath := "D:\\Mir\\test\\aaa\\export\\241121\\in_241121\\"
	//basePath := "D:\\Mir\\test\\aaa\\export\\241124\\in_1124\\"
	basePath := "D:\\Mir\\test\\aaa\\export\\1202\\"
	for i := 1; i <= 383; i++ {
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
		methods.JsonToExcel(string(data), "第二批数据", strconv.Itoa(i))
	}
}

// 打印出第二张表相较于第一张表多出来的数据
func (u UserController) ExportAdd(c *gin.Context) {

	// 1. 读取文件: D:\Mir\test\aaa\export\第一批数据.xlsx 做为第一批数据
	file1, err := xlsx.OpenFile("D:\\Mir\\test\\aaa\\export\\第一批数据.xlsx")
	if err != nil {
		log.Println("读取第一批数据文件出错:", err)
		return
	}
	cardIndex := 0                       // 身份证号码所在的列
	cardMap := make(map[string]struct{}) // 创建了一个键类型为 string ，值类型为空结构体 struct{} 的 map, 表示只关心键的存在，而不需要为每个键关联具体的值
	for i, row := range file1.Sheets[0].Rows {
		if i == 0 {
			// 当前是在第一行（表头） 将列名为身份证号的存在cardIndex中
			for ci, cell := range row.Cells {
				if cell.Value == "身份证号" {
					cardIndex = ci
				}
			}
		} else {
			if len(row.Cells) > 0 {
				value := row.Cells[cardIndex].Value
				// 向cardMap中添加元素
				// struct{} 表示一个空的结构体类型，后面的 {} 是创建并初始化这个空结构体类型的一个值。
				// 使用空结构体作为值主要是因为它不占用内存空间，仅用于表示某个键的存在与否。这种方式简单且高效，适用于只需要标记某个键是否存在，而不需要为每个键存储额外信息的场景。
				cardMap[value] = struct{}{}
			}
		}
	}

	// 2. 读取文件: D:\Mir\test\aaa\export\第二批数据.xlsx 做为第二批数据
	file2, err := xlsx.OpenFile("D:\\Mir\\test\\aaa\\export\\第二批数据.xlsx")
	if err != nil {
		log.Println("读取第二批数据文件出错:", err)
		return
	}

	cardIndex2 := 0            // 身份证号码所在的列
	var compareData [][]string // 对比出来的数据放在这里 二维字符串切片因为需要存储多行多列的数据结构

	for i, row := range file2.Sheets[0].Rows {
		if i == 0 {
			// 当前是在第一行（表头） 将列名为身份证号的存在cardIndex中
			for ci, cell := range row.Cells {
				if cell.Value == "身份证号" {
					cardIndex2 = ci
				}
			}
		} else {
			if len(row.Cells) > 0 {
				// 判断第一批数据表中是否存在当前身份证号 不存在 将其添加在compareData
				target := row.Cells[cardIndex2].Value
				_, exists := cardMap[target]
				if !exists {
					// 从每个 Cell 中获取值并转换为字符串后再添加
					var rowData []string
					for _, cell := range row.Cells {
						rowData = append(rowData, cell.Value)
					}
					compareData = append(compareData, rowData)
				}
			}
		}
	}

	// 把对比文件的表头也放进来
	headers := []string{
		"出生日期", "年龄", "社保卡号", "医保卡号", "zfkh", "cardType", "cardNo", "idcard", "签约状态", "签约时间",
		"身份证号", "军官证号", "学生证号", "居住证号", "居民姓名", "居民性别", "联系电话", "手机号码", "居住地址", "配送地址", "医生执业证号",
		"医生编码", "签约医生", "医生电话", "社区机构编码", "社区机构名称", "qx", "区级机构编码", "区级机构名称", "市级机构编码", "市级机构名称",
		"操作机构编码", "操作机构名称", "操作人员编码", "操作人员姓名", "jlsj", "bgsj", "account", "appid", "authtype",
		"authpath", "fromDate", "toDate", "isauth", "jlsjc", "bgsjc", "fromDatec", "到期时间",
		"人群标签", "renewable", "provinceCode", "cityCode", "districtCode", "townCode", "villageCode",
		"proxySfzh", "isproxy", "iscaface", "qhsj", "proxyName", "pgid", "orgAble", "jgbgsj",
		"qyxxly", "gzr", "gzrqt", "gzsj", "id", "fjid", "xysj", "xysjc", "区", "居委",
		"lxrList", "swrq", "swbgdw", "swbgrq", "jslxfs", "swlxrxm", "swlxrdh", "rksj", "phone",
		"keyGroup", "jyrq", "xygzfs",
	}
	compareData = append([][]string{headers}, compareData...)

	// 3. 将对比出的数据放在 新建的表格 对比文件 中。新建的表格地址如下：D:\Mir\test\aaa\export\对比文件.xlsx
	newFile := xlsx.NewFile()
	sheet, err := newFile.AddSheet("对比结果")
	if err != nil {
		fmt.Println("创建新表格出错:", err)
		return
	}

	for _, rowData := range compareData {
		row := sheet.AddRow()
		for _, data := range rowData {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = newFile.Save("D:\\Mir\\test\\aaa\\export\\对比文件.xlsx")
	if err != nil {
		fmt.Println("保存对比文件出错:", err)
		return
	}

}
