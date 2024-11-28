
// 241127 读取同级目录下 第一批数据表和第二批数据表 对比出其中第二批数据表中存在 && 第一批数据表中不存在的数据
console.log('-----------开始----------')
const XLSX = require('xlsx');

// 导出表格
const exportExcel = (data, fileName) => {
    // 创建一个新的工作簿
    let workbook = XLSX.utils.book_new()
    // 将数据转换为工作表
    let worksheet = XLSX.utils.json_to_sheet(data)
    // 将工作表添加到工作簿中，命名为 'Sheet1'
    XLSX.utils.book_append_sheet(workbook, worksheet, 'Sheet1')
    // 将工作簿写入为 CSV 文件`
    XLSX.writeFile(workbook, `${fileName}`)
};

// 读取 已有的第一批数据表
const workbook = XLSX.readFile('第一批数据.xlsx');
// 获取第一个工作表
const sheetName = workbook.SheetNames[0];
const sheet = workbook.Sheets[sheetName];
// 将数据库导出表数据转换为 JSON
const data = XLSX.utils.sheet_to_json(sheet);

// 读取 已有的第二批数据表
const workbook2 = XLSX.readFile('第二批数据.xlsx');
// 获取第一个工作表
const sheetName2 = workbook2.SheetNames[0];
const sheet2 = workbook2.Sheets[sheetName2];
// 将数据库导出表数据转换为 JSON
const data2 = XLSX.utils.sheet_to_json(sheet2);

// 将第一批数据中的身份证号都存起来用于做比较
cardArr = []
data.forEach(item => {
    cardArr.push(item["身份证号"])
})

let result = []
data2.forEach(item2 => {
    // 如果第二批数据中含有 不存在于第一批的数据 那么将其 放入最终的列表
    if (!cardArr.includes(item2["身份证号"])) {
        result.push(item2)
    }
})
console.log("对比出", result.length, "条数据")

exportExcel(result, "第二批中新增的数据.xlsx")

console.log('-------------结束---------------')