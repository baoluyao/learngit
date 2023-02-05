package ali

import (
	"noticeservice/pkg/app"
	"strings"

	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

// 创建短信模板
// if _, _, err := AddSmsTemplate("机器故障", "${name}您好!机器人${robot}在${time}有故障产生，请及时排查。", "链接：http://www.ecovacs-c.com 使用场景：当机器人出现异常时将短信通知作为其中一种方式，可及时有效帮助机器人排除故障进而继续工作"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

// 编辑短信模板
// if code, msg, err := ModifySmsTemplate("机器故障", ":SMS_241355231", "${name}您好!机器人${robot}在${time}出故障了，请及时排查。", "链接：http://www.ecovacs-c.com 使用场景：当机器人出现异常时将短信通知作为其中一种方式，可及时有效帮助机器人排除故障进而继续工作"); err != nil {
// 	fmt.Println("error:", err.Error())
// } else {
// 	if code != "" {
// 		fmt.Println("code:", code, "msg:", msg)
// 	}
// }

// 查询指定模板详细
// if status, _, _, err := QuerySmsTemplate(":SMS_241355231"); err != nil {
// 	fmt.Println("error:", err.Error())
// } else {
// 	fmt.Println("status:", status)
// }

// 删除指定模板
// if _, _, err := DeleteSmsTemplate(":SMS_241355231"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

// 查询模板列表
// if temlates, _, _, err := QuerySmsTemplateList(); err != nil {
// 	fmt.Println("error:", err.Error())
// } else {
// 	fmt.Println("temlates:", temlates)
// }

// 发送短信
func (ali *AliClient) SendSms(phoneNumbers []string, signName string, templateCode string, templateParam string) (code string, codemsg string, _err error) {

	_phoneNumbers := strings.Join(phoneNumbers, ",")
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(_phoneNumbers), //收信人电话号码  可填多个
		SignName:      tea.String(signName),      //短信签名
		TemplateCode:  tea.String(templateCode),  //模板编码
		TemplateParam: tea.String(templateParam), //模板参数变量
	}
	// 执行操作
	result, _err := ali.dysmsClient.SendSms(sendSmsRequest)
	if _err != nil {
		return "", "", _err
	}

	code = tea.StringValue(result.Body.Code)
	codemsg = tea.StringValue(result.Body.Message)

	return code, codemsg, nil
}

//批量发送短信
func (ali *AliClient) SendBatchSms(phoneNumberJson string, signNameJson string, TemplateCode string, templateParamJson string) (code string, codemsg string, _err error) {
	sendBatchSmsRequest := &dysmsapi20170525.SendBatchSmsRequest{
		PhoneNumberJson:   tea.String(phoneNumberJson),
		SignNameJson:      tea.String(signNameJson),
		TemplateCode:      tea.String(TemplateCode),
		TemplateParamJson: tea.String(templateParamJson),
		//SmsUpExtendCodeJson: tea.String(smsUpExtendCodeJson),
	}
	result, _err := ali.dysmsClient.SendBatchSms(sendBatchSmsRequest)
	if _err != nil {
		app.WriteWorkLog("SendBatchSms fail, error:%v", _err)
		return "", "", _err
	}
	app.WriteWorkLog("SendBatchSms success, params:%v", sendBatchSmsRequest)

	code = tea.StringValue(result.Body.Code)
	codemsg = tea.StringValue(result.Body.Message)

	return code, codemsg, nil
}

// 查询短信模板列表
func (ali *AliClient) QuerySmsTemplateList() (templates []interface{}, code, msg string, _err error) {

	querySmsTemplateListRequest := &dysmsapi20170525.QuerySmsTemplateListRequest{}
	// 执行操作
	result, _err := ali.dysmsClient.QuerySmsTemplateList(querySmsTemplateListRequest)
	if _err != nil {
		return nil, "", "", _err
	}

	code = tea.StringValue(result.Body.Code)
	msg = tea.StringValue(result.Body.Message)

	for _, _template := range result.Body.SmsTemplateList {
		template := map[string]interface{}{
			"templateCode":    tea.StringValue(_template.TemplateCode),
			"auditStatus":     tea.StringValue(_template.AuditStatus),
			"templateContent": tea.StringValue(_template.TemplateContent),
			"templateName":    tea.StringValue(_template.TemplateName),
			"templateType":    tea.Int32Value(_template.TemplateType),
			"orderId":         tea.StringValue(_template.OrderId),
			"createDate":      tea.StringValue(_template.CreateDate),
			"reason":          _template.Reason,
		}

		templates = append(templates, template)
	}

	return templates, code, msg, nil
}

// 创建短信模板
func (ali *AliClient) AddSmsTemplate(templateName, templateContent, remark string) (code, msg string, _err error) {

	addSmsTemplateRequest := &dysmsapi20170525.AddSmsTemplateRequest{
		TemplateType:    tea.Int32(1),
		TemplateName:    tea.String(templateName),
		TemplateContent: tea.String(templateContent),
		Remark:          tea.String(remark),
	}
	// 执行操作
	result, _err := ali.dysmsClient.AddSmsTemplate(addSmsTemplateRequest)
	if _err != nil {
		return "", "", _err
	}

	code = tea.StringValue(result.Body.Code)
	msg = tea.StringValue(result.Body.Message)

	return code, msg, nil
}

// 编辑短信模板
func (ali *AliClient) ModifySmsTemplate(templateName, templateCode, templateContent, remark string) (code string, codemsg string, _err error) {

	modifySmsTemplateRequest := &dysmsapi20170525.ModifySmsTemplateRequest{
		TemplateType:    tea.Int32(1),
		TemplateName:    tea.String(templateName),
		TemplateCode:    tea.String(templateCode),
		TemplateContent: tea.String(templateContent),
		Remark:          tea.String(remark),
	}
	// 执行操作
	result, _err := ali.dysmsClient.ModifySmsTemplate(modifySmsTemplateRequest)
	if _err != nil {
		return "", "", _err
	}

	code = tea.StringValue(result.Body.Code)
	codemsg = tea.StringValue(result.Body.Message)

	return code, codemsg, nil
}

// 查询模板审核状态
func (ali *AliClient) QuerySmsTemplate(templateCode string) (status int, code, msg string, _err error) {

	querySmsTemplateRequest := &dysmsapi20170525.QuerySmsTemplateRequest{
		TemplateCode: tea.String(templateCode),
	}
	// 执行操作
	result, _err := ali.dysmsClient.QuerySmsTemplate(querySmsTemplateRequest)
	if _err != nil {
		return -1, "", "", _err
	}

	code = tea.StringValue(result.Body.Code)
	msg = tea.StringValue(result.Body.Message)

	status = int(tea.Int32Value(result.Body.TemplateStatus))

	return status, code, msg, nil
}

// 删除指定的模板
func (ali *AliClient) DeleteSmsTemplate(templateCode string) (code, msg string, _err error) {

	deleteSmsTemplateRequest := &dysmsapi20170525.DeleteSmsTemplateRequest{
		TemplateCode: tea.String(templateCode),
	}
	// 执行操作
	result, _err := ali.dysmsClient.DeleteSmsTemplate(deleteSmsTemplateRequest)
	if _err != nil {
		return "", "", _err
	}

	code = tea.StringValue(result.Body.Code)
	msg = tea.StringValue(result.Body.Message)

	return code, msg, nil
}
