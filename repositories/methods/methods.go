package methods

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
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
