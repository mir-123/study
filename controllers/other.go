package controllers

import (
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
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

// 获取对比出来的医生数据
func (u UserController) GetCompare(c *gin.Context) {

	path := "D:\\Mir\\test\\aaa\\other_temp\\info(1).csv"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建CSV Reader对象
	reader := csv.NewReader(file)

	// 逐行读取数据并处理
	cot := 0
	for {
		cot++
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("×", record[0], len(record))
			fmt.Println(err)
			return
		}

		if len(record) != 64 {
			fmt.Println(record)
			return
		}

		if cot == 1 {
			continue
		}

		_, err = strconv.Atoi(record[0])
		if err != nil {
			fmt.Println(err, cot, record)
			return
		}
	}

	// 可以实现
	////csvPath := "D:\\Mir\\test\\aaa\\other_temp\\test.csv"
	//csvPath := "D:\\Mir\\test\\aaa\\other_temp\\info(1).csv"
	//jsonData, err := csvToJSON(csvPath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(jsonData))
}

func csvToJSON(csvPath string) ([]byte, error) {
	type Record struct {
		Value string `json:"value"`
	}

	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var data []Record
	for _, record := range records {
		if len(record) > 0 {
			data = append(data, Record{Value: record[0]})
		}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
