package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
	"gopkg.in/gomail.v2"
)

type Codes struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var (
	Active = NewCode(0, "正常")
	//Enable  = NewCode(1, "启用")
	Disable = NewCode(1, "停用")
	Delete  = NewCode(2, "删除")
	Expired = NewCode(3, "失效")
)
var codes = map[int]string{}

func NewCode(code int, msg string) *Codes {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("Code码%d已经存在,请更换一个", code))
	}
	codes[code] = msg
	return &Codes{code: code, msg: msg}
}
func (e *Codes) Code() int {
	return e.code
}

func Validuuid(uid uuid.UUID) bool {
	ss := uid.String()
	if ss == "00000000-0000-0000-0000-000000000000" {
		return false
	} else {
		_, err := uuid.Parse(ss)
		if err != nil {
			return false
		}
	}
	return true
}

// Contains 数组是否包含某元素
func Contains(slice []string, s string) int {
	for index, value := range slice {
		if value == s {
			return index
		}
	}
	return -1
}

func MapToStruct(data interface{}, v interface{}) error {
	bin, _ := json.Marshal(data)
	return json.Unmarshal(bin, v)
}
func ToJsonBin(obj interface{}) []byte {
	bin, _ := json.Marshal(obj)
	return bin
}

func EncodeURI(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}
func DataToFile(w http.ResponseWriter, r *http.Request, fileName string, data []byte) {

	fileName = fmt.Sprintf(`attachment; filename=%s`, EncodeURI(fileName))
	w.Header().Add("Content-Disposition", fileName)
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Content-Transfer-Encoding", "binary")

	content := bytes.NewReader(data)
	http.ServeContent(w, r, fileName, time.Now(), content)
}

//导出excel
func DataToExcel(w http.ResponseWriter, r *http.Request, titleList []map[string]string, dataList []map[string]interface{}, fileName string) {
	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleRow := sheet.AddRow()
	// for _, v := range titleList {
	for v := 0; v < len(titleList); v++ {
		cell := titleRow.AddCell()
		item := titleList[v]
		for _, name := range item {
			cell.Value = name
			cell.GetStyle().Font.Bold = true
		}
		// cell.GetStyle().Font.Color = "00FF0000"
	}
	// 插入内容
	//range不能按顺序输出
	for v := 0; v < len(dataList); v++ {
		row := sheet.AddRow()
		for _, item := range titleList {
			for key, _ := range item {
				cell := row.AddCell()
				cell.SetValue(dataList[v][key])
			}
		}
	}
	fileName = fmt.Sprintf("%s.xlsx", EncodeURI(fileName))
	//_ = file.Save(fileName)
	// w.Header().Add("filename", fileName)
	// fileName = fmt.Sprintf(`attachment; filename=%s`, EncodeURI(fileName))
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	http.ServeContent(w, r, fileName, time.Now(), content)
}

func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码

	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",
	//  "port": "465",
	//}

	mailConn := map[string]string{
		"user": "xxx@sample.cn",
		"pass": "r4r3St*****7a7Uk",
		"host": "smtp.exmail.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "XX官方"))
	/*这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])*/
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

// func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *onsmqtt20200420.Client, _err error) {
// 	config := &openapi.Config{
// 		// 您的AccessKey ID
// 		AccessKeyId: accessKeyId,
// 		// 您的AccessKey Secret
// 		AccessKeySecret: accessKeySecret,
// 	}
// 	// 访问的域名
// 	config.Endpoint = tea.String(global.AliMQTTSetting.EndPoit)
// 	_result = &onsmqtt20200420.Client{}
// 	_result, _err = onsmqtt20200420.NewClient(config)
// 	return _result, _err
// }

// func GetAPKInfo(path string) (string, string, error) {
// 	resp, err := http.Get(path)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	defer resp.Body.Close()
// 	thebytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	reader := bytes.NewReader(thebytes)
// 	app, err := apk.OpenZipReader(reader, resp.ContentLength)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	//app, _ = apk.SetFile(app, reader)
// 	return app.PackageName(), app.Manifest().VersionName.MustString(), nil //
// }
