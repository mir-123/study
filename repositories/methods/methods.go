package methods

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"os"
	"strconv"
)

// 根据 json 文件导出 excel 表格
type DataItem struct {
	Csrq         string `json:"csrq"`
	Age          uint8  `json:"age"`
	Sbkh         string `json:"sbkh"`
	Ybkh         string `json:"ybkh"`
	Zfkh         string `json:"zfkh"`
	CardType     string `json:"cardType"`
	CardNo       string `json:"cardNo"`
	Idcard       string `json:"idcard"`
	Qystate      string `json:"qystate"`
	Qysj         string `json:"qysj"`
	Sfzh         string `json:"sfzh"`
	Jgzh         string `json:"jgzh"`
	Xszh         string `json:"xszh"`
	Jzzh         string `json:"jzzh"`
	Jmxm         string `json:"jmxm"`
	Jmxb         string `json:"jmxb"`
	Lxdh         string `json:"lxdh"`
	Sjhm         string `json:"sjhm"`
	Jjdz         string `json:"jjdz"`
	Psdz         string `json:"psdz"`
	Yszyzh       string `json:"yszyzh"`
	GpId         string `json:"gpId"`
	GpMc         string `json:"gpMc"`
	GpDh         string `json:"gpDh"`
	Sqjgbm       string `json:"sqjgbm"`
	Sqjgmc       string `json:"sqjgmc"`
	Qx           string `json:"qx"`
	Qjjgbm       string `json:"qjjgbm"`
	Qjjgmc       string `json:"qjjgmc"`
	Sjjgbm       string `json:"sjjgbm"`
	Sjjgmc       string `json:"sjjgmc"`
	Czjgbm       string `json:"czjgbm"`
	Czjgmc       string `json:"czjgmc"`
	Czrybm       string `json:"czrybm"`
	Czryxm       string `json:"czryxm"`
	Jlsj         string `json:"jlsj"`
	Bgsj         string `json:"bgsj"`
	Account      string `json:"account"`
	Appid        string `json:"appid"`
	Authtype     string `json:"authtype"`
	Authpath     string `json:"authpath"`
	FromDate     string `json:"fromDate"`
	ToDate       string `json:"toDate"`
	Isauth       string `json:"isauth"`
	Jlsjc        string `json:"jlsjc"`
	Bgsjc        string `json:"bgsjc"`
	FromDatec    string `json:"fromDatec"`
	ToDatec      string `json:"toDatec"`
	Color        string `json:"color"`
	Renewable    string `json:"renewable"`
	ProvinceCode string `json:"provinceCode"`
	CityCode     string `json:"cityCode"`
	DistrictCode string `json:"districtCode"`
	TownCode     string `json:"townCode"`
	VillageCode  string `json:"villageCode"`
	ProxySfzh    string `json:"proxySfzh"`
	Isproxy      string `json:"isproxy"`
	Iscaface     string `json:"iscaface"`
	Qhsj         string `json:"qhsj"`
	ProxyName    string `json:"proxyName"`
	Pgid         string `json:"pgid"`
	OrgAble      string `json:"orgAble"`
	Jgbgsj       string `json:"jgbgsj"`
	Qyxxly       string `json:"qyxxly"`
	Gzr          string `json:"gzr"`
	Gzrqt        string `json:"gzrqt"`
	Gzsj         string `json:"gzsj"`
	ID           string `json:"id"`
	Fjid         string `json:"fjid"`
	Xysj         string `json:"xysj"`
	Xysjc        string `json:"xysjc"`
	Qxmc         string `json:"qxmc"`
	Yyjc         string `json:"yyjc"`
	LxrList      string `json:"lxrList"`
	Swrq         string `json:"swrq"`
	Swbgdw       string `json:"swbgdw"`
	Swbgrq       string `json:"swbgrq"`
	Jslxfs       string `json:"jslxfs"`
	Swlxrxm      string `json:"swlxrxm"`
	Swlxrdh      string `json:"swlxrdh"`
	Rksj         string `json:"rksj"`
	Phone        string `json:"phone"`
	KeyGroup     string `json:"keyGroup"`
	Jyrq         string `json:"jyrq"`
	Xygzfs       string `json:"xygzfs"`
}

func JsonToExcel(jsonData string, filename string, name string) {
	// 解析 JSON 数据
	var data struct {
		PageSize int        `json:"pageSize"`
		Total    int        `json:"total"`
		Data     []DataItem `json:"data"`
		CurrPage int        `json:"currPage"`
		Pages    int        `json:"pages"`
	}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("解析 JSON 数据出错:", err)
		return
	}

	// 打开或创建指定的 Excel 文件
	file, err := xlsx.OpenFile("D:\\Mir\\test\\aaa\\export\\" + filename + ".xlsx")
	if err != nil {
		// 如果文件不存在，则创建一个新的
		file = xlsx.NewFile()
	}
	sheet, ok := file.Sheet["Sheet1"]
	if !ok {
		sheet, err = file.AddSheet("Sheet1")
		if err != nil {
			fmt.Println("创建 Excel 工作表出错:", err)
			return
		}
	}

	// 写入表头
	if sheet.Rows == nil {
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
		row := sheet.AddRow()
		for _, header := range headers {
			cell := row.AddCell()
			cell.Value = header
		}
	}

	// 写入数据
	for _, item := range data.Data {
		row := sheet.AddRow()
		row.AddCell().Value = item.Csrq
		row.AddCell().Value = fmt.Sprintf("%d", item.Age)
		row.AddCell().Value = item.Sbkh
		row.AddCell().Value = item.Ybkh
		row.AddCell().Value = item.Zfkh
		row.AddCell().Value = item.CardType
		row.AddCell().Value = item.CardNo
		row.AddCell().Value = item.Idcard
		row.AddCell().Value = item.Qystate
		row.AddCell().Value = item.Qysj
		row.AddCell().Value = item.Sfzh
		row.AddCell().Value = item.Jgzh
		row.AddCell().Value = item.Xszh
		row.AddCell().Value = item.Jzzh
		row.AddCell().Value = item.Jmxm
		if item.Jmxb == "1" {
			row.AddCell().Value = "男"
		} else if item.Jmxb == "2" {
			row.AddCell().Value = "女"
		} else {
			row.AddCell().Value = item.Jmxb
		}
		row.AddCell().Value = item.Lxdh
		row.AddCell().Value = item.Sjhm
		row.AddCell().Value = item.Jjdz
		row.AddCell().Value = item.Psdz
		row.AddCell().Value = item.Yszyzh
		row.AddCell().Value = item.GpId
		row.AddCell().Value = item.GpMc
		row.AddCell().Value = item.GpDh
		row.AddCell().Value = item.Sqjgbm
		row.AddCell().Value = item.Sqjgmc
		row.AddCell().Value = item.Qx
		row.AddCell().Value = item.Qjjgbm
		row.AddCell().Value = item.Qjjgmc
		row.AddCell().Value = item.Sjjgbm
		row.AddCell().Value = item.Sjjgmc
		row.AddCell().Value = item.Czjgbm
		row.AddCell().Value = item.Czjgmc
		row.AddCell().Value = item.Czrybm
		row.AddCell().Value = item.Czryxm
		row.AddCell().Value = item.Jlsj
		row.AddCell().Value = item.Bgsj
		row.AddCell().Value = item.Account
		row.AddCell().Value = item.Appid
		row.AddCell().Value = item.Authtype
		row.AddCell().Value = item.Authpath
		row.AddCell().Value = item.FromDate
		row.AddCell().Value = item.ToDate
		row.AddCell().Value = item.Isauth
		row.AddCell().Value = item.Jlsjc
		row.AddCell().Value = item.Bgsjc
		row.AddCell().Value = item.FromDatec
		row.AddCell().Value = item.ToDatec
		if item.Color == "red" {
			row.AddCell().Value = "红色"
		} else if item.Color == "yellow" {
			row.AddCell().Value = "黄色"
		} else if item.Color == "green" {
			row.AddCell().Value = "绿色"
		} else {
			row.AddCell().Value = item.Color
		}
		row.AddCell().Value = item.Renewable
		row.AddCell().Value = item.ProvinceCode
		row.AddCell().Value = item.CityCode
		row.AddCell().Value = item.DistrictCode
		row.AddCell().Value = item.TownCode
		row.AddCell().Value = item.VillageCode
		row.AddCell().Value = item.ProxySfzh
		row.AddCell().Value = item.Isproxy
		row.AddCell().Value = item.Iscaface
		row.AddCell().Value = item.Qhsj
		row.AddCell().Value = item.ProxyName
		row.AddCell().Value = item.Pgid
		row.AddCell().Value = item.OrgAble
		row.AddCell().Value = item.Jgbgsj
		row.AddCell().Value = item.Qyxxly
		row.AddCell().Value = item.Gzr
		row.AddCell().Value = item.Gzrqt
		row.AddCell().Value = item.Gzsj
		row.AddCell().Value = item.ID
		row.AddCell().Value = item.Fjid
		row.AddCell().Value = item.Xysj
		row.AddCell().Value = item.Xysjc
		row.AddCell().Value = item.Qxmc
		row.AddCell().Value = item.Yyjc
		row.AddCell().Value = item.LxrList
		row.AddCell().Value = item.Swrq
		row.AddCell().Value = item.Swbgdw
		row.AddCell().Value = item.Swbgrq
		row.AddCell().Value = item.Jslxfs
		row.AddCell().Value = item.Swlxrxm
		row.AddCell().Value = item.Swlxrdh
		row.AddCell().Value = item.Rksj
		row.AddCell().Value = item.Phone
		row.AddCell().Value = item.KeyGroup
		row.AddCell().Value = item.Jyrq
		row.AddCell().Value = item.Xygzfs
	}

	// 保存 Excel 文件
	err = file.Save("D:\\Mir\\test\\aaa\\export\\" + filename + ".xlsx")
	if err != nil {
		fmt.Println("保存 Excel 文件出错:", err)
		return
	}
	fmt.Printf("转换成功，%s\n", name)
}

func ReadCsv(path string) {
	// 打开CSV文件
	// 0 id
	// 1 uuid
	// 2 data_from
	// 3 serial_code
	// 4 org_area_code
	// 5 unit_id
	// 6 institution_indentity
	// 7 account
	// 8 password
	// 9 name
	// 10 name_en
	// 11 gender
	// 12 birthday
	// 13 photo_path
	// 14 indentity_code
	// 15 nationality
	// 16 native_place
	// 17 politics_position
	// 18 working_time
	// 19 top_education
	// 20 top_degree
	// 21 graduate_school
	// 22 graduate_date
	// 23 office_phone
	// 24 office_fax_no
	// 25 mobile
	// 26 email
	// 27 address
	// 28 credentials_no
	// 29 date_of_get_credential
	// 30 license_no
	// 31 date_of_get_license
	// 32 department
	// 33 dep_code
	// 34 grade
	// 35 position
	// 36 date_of_hold_post
	// 37 job_title
	// 38 date_of_get_job_title
	// 39 working_status
	// 40 work_years
	// 41 practice_scope
	// 42 practice_sort
	// 43 practice_specialty
	// 44 summary
	// 45 remark
	// 46 status
	// 47 doctor_profession
	// 48 profession_edit_uid
	// 49 certificate_type
	// 50 certificate_code
	// 51 create_by
	// 52 create_on
	// 53 update_by
	// 54 update_on
	// 55 login_count
	// 56 last_login
	// 57 area
	// 58 province
	// 59 city
	// 60 district
	// 61 is_push
	// 62 is_del
	// 63 is_duplicate
	// 源文件行数: 1633967
	file, err := os.Open(path) // data.csv为要读取的CSV文件名
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// ww, err := os.OpenFile("temp.csv", os.O_WRONLY|os.O_CREATE, 0777) // data.csv为要读取的CSV文件名
	// if err != nil {
	//         panic(err)
	// }
	// w := csv.NewWriter(ww)
	// err = w.Write([]string{"1", "2,\"\"2,2", "3\"", "4", "5.0", "6\n6"})
	// if err != nil {
	//         fmt.Println(err)
	//         return
	// }
	// w.Flush()
	// ww.Close()
	// if time.Now().Unix() > 1 {
	//         return
	// }

	// buf := bufio.NewReader(file)
	// c := 0
	// lastID := 0
	// for {

	//         // ReadString reads until the first occurrence of delim in the input,
	//         // returning a string containing the data up to and including the delimiter.
	//         line, err := buf.ReadString('\n')
	//         c++
	//         if err != nil {
	//                 if err == io.EOF {
	//                         fmt.Println(string(line))
	//                         break
	//                 }
	//                 fmt.Println(err)
	//                 return
	//         }
	//         if c == 1 {
	//                 continue
	//         }

	//         lineS := string(line)
	//         i := strings.Index(lineS, ",")
	//         if i == -1 {
	//                 fmt.Println(c, i)
	//                 return
	//         }
	//         id, err := strconv.Atoi(lineS[:i])
	//         if err != nil {
	//                 fmt.Println(err, c, i)
	//                 return
	//         }
	//         if id < lastID {
	//                 panic(id)
	//         }
	//         lastID = id

	//         // if !prefix {
	//         // }
	//         // if c == 1633966 {
	//         // fmt.Println(string(line), prefix)
	//         // }
	//         // params := strings.Split(string(line), ",")
	//         // if len(params) != 64 {
	//         //         fmt.Println(string(line), len(params))
	//         //         return
	//         // }
	// }
	// fmt.Println(c)
	// if c > 2 {
	//         return
	// }

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
			// fmt.Println(err)
			// return
		}

		// 输出每行记录内容
		// fmt.Println(record[0])
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
		// if id < lastID {
		//         fmt.Println(id, lastID, cot)
		//         return
		// }
		//telephone[record[25]]++
		// if record[25] == "好" {
		//         fmt.Println(path, record)
		// }
	}
}
