package ali

import (
	"strconv"

	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	"github.com/alibabacloud-go/tea/tea"
)

//创建域名
func (ali *AliClient) CreateDomain(domain string) (_err error) {

	createDomainRequest := &dm20151123.CreateDomainRequest{
		DomainName: tea.String(domain),
	}
	//执行操作
	_, _err = ali.dmClient.CreateDomain(createDomainRequest)
	if _err != nil {
		return _err
	}
	return _err
}

//发送单封邮件
// if err := SingleSendMail("[测试邮件]机器人异常了，请及时排查。", "zhongkun.zhou@ecovacs.com,luyao.bao@ecovacs.com", "test","public@marketing.ecovacs-c.com"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

// if err := SingleSendMail("[测试邮件]机器人异常了，请及时排查。", "luyao.bao@ecovacs.com", "test", "public@marketing.ecovacs-c.com"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

//批量发送邮件
// err := BatchSendMail("ecovscsBiz", "public@marketing.ecovacs-c.com", "ecoeco")
// if err != nil {
// 	fmt.Println("error:", err.Error())
// }

//创建域名
// if err := CreateDomain("dev.ecovacs-c.com"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

//获取域名列表
// if domains, err := QueryDomainByParam(); err != nil {
// 	fmt.Println("error:", err.Error())
// } else {
// 	fmt.Println("domians:", domains)
// }

// 删除域名
// if err := DeleteDomain("345136"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

//查询发信地址列表
// if mas, err := QueryMailAddressByParam(); err != nil {
// 	fmt.Println("error:", err.Error())
// } else {
// 	fmt.Println("mas:", mas)
// }

//创建发信地址
// if err := CreateMailAddress("dev@marketing.ecovacs-c.com", "trigger"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

//删除发信地址
// if err := DeleteMailAddress("283571"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

//查询收信人列表
// if recvs, err := QueryReceiverByParam(); err != nil {
// 	fmt.Println("error:", err.Error())
// } else {
// 	fmt.Println("recvs:", recvs)
// }

//查询指定收信人详细
// if users, err := QueryReceiverDetail("ad0ba0dc502e7bb8ba832db24081d089"); err != nil {
// 	fmt.Println("error:", err.Error())
// } else {
// 	fmt.Println("users:", users)
// }

//创建收信列表
// if err := CreateReceiver("ecoeco", "eco@126.com"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

//删除收信列表
// if err := DeleteReceiver("ad0ba0dc502e7bb8ba832db24081d089"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

//添加收信人到收信列表
// receiver := map[string]interface{}{
// 	"b": "1989/9/4",
// 	"e": "luyao.bao@ecovacs.com",
// 	"g": "先生",
// 	"m": "17365379289",
// 	"n": "坤坤",
// 	"u": "周中坤",
// }
// var receivers []interface{} = make([]interface{}, 0)
// receivers = append(receivers, receiver)
// rawByte, err := json.Marshal(receivers)
// if err != nil {
// 	fmt.Println("error:", err.Error())
// } else {
// 	if err := SaveReceiverDetail("ad0ba0dc502e7bb8ba832db24081d089", string(rawByte)); err != nil {
// 		fmt.Println("error:", err.Error())
// 	}
// }

//从收信列表删除指定收信人
// if err := DeleteReceiverDetail("ad0ba0dc502e7bb8ba832db24081d089", "luyao.bao@ecovacs.com"); err != nil {
// 	fmt.Println("error:", err.Error())
// }

//发送单封邮件
func (ali *AliClient) SingleSendMail(content string, toEmail string, subject string, fromEmail string) (_err error) {

	singleSendMailRequest := &dm20151123.SingleSendMailRequest{
		AccountName:    tea.String(fromEmail), //管理控制台中配置的发信地址
		AddressType:    tea.Int32(1),          //地址类型。取值：0：为随机账号 1：为发信地址
		ToAddress:      tea.String(toEmail),   //目标地址
		Subject:        tea.String(subject),   //邮件主题
		TextBody:       tea.String(content),   //邮件正文
		ReplyToAddress: tea.Bool(false),       //回信地址
	}
	//执行操作
	_, _err = ali.dmClient.SingleSendMail(singleSendMailRequest)
	if _err != nil {
		return _err
	}

	return nil
}

//批量发送邮件
func (ali *AliClient) BatchSendMail(templateName, fromEmail, receiversName string) (_err error) {
	batchSendMailRequest := &dm20151123.BatchSendMailRequest{
		TemplateName:  tea.String(templateName),  // 模板名  预先创建且通过审核的模板名称
		AccountName:   tea.String(fromEmail),     // 发信地址  管理控制台中配置的发信地址
		ReceiversName: tea.String(receiversName), // 收信人列表名称  预先创建且上传了收件人的收件列表名称
		AddressType:   tea.Int32(0),              // 账号类型 0：随机地址 1：发信地址
	}
	//执行操作
	_, _err = ali.dmClient.BatchSendMail(batchSendMailRequest)
	if _err != nil {
		return _err
	}
	return _err
}

///获取域名列表
func (ali *AliClient) QueryDomainByParam() (domains []interface{}, _err error) {

	queryDomainByParamRequest := &dm20151123.QueryDomainByParamRequest{}
	//执行操作
	resp, _err := ali.dmClient.QueryDomainByParam(queryDomainByParamRequest)
	if _err != nil {
		return nil, _err
	}

	for _, _domain := range resp.Body.Data.Domain {
		domain := map[string]interface{}{
			"id":         tea.StringValue(_domain.DomainId),
			"name":       tea.StringValue(_domain.DomainName),
			"status":     tea.StringValue(_domain.DomainStatus),
			"createTime": tea.StringValue(_domain.CreateTime),
		}
		domains = append(domains, domain)
	}

	return domains, nil
}

//删除域名
func (ali *AliClient) DeleteDomain(domainId string) (_err error) {

	var _domainId int32 = 0
	id, _err := strconv.Atoi(domainId)
	_domainId = int32(id)
	if _err != nil {
		return _err
	}
	deleteDomainRequest := &dm20151123.DeleteDomainRequest{
		DomainId: tea.Int32(_domainId),
	}
	//执行操作
	_, _err = ali.dmClient.DeleteDomain(deleteDomainRequest)
	if _err != nil {
		return _err
	}
	return _err
}

//查询发信地址
func (ali *AliClient) QueryMailAddressByParam() (mas []interface{}, _err error) {

	queryMailAddressByParamRequest := &dm20151123.QueryMailAddressByParamRequest{}
	//执行操作
	address, _err := ali.dmClient.QueryMailAddressByParam(queryMailAddressByParamRequest)
	if _err != nil {
		return nil, _err
	}

	for _, _addr := range address.Body.Data.MailAddress {
		addr := map[string]interface{}{
			"id":            tea.StringValue(_addr.MailAddressId),
			"sendtype":      tea.StringValue(_addr.Sendtype),
			"nmae":          tea.StringValue(_addr.AccountName),
			"statu":         tea.StringValue(_addr.AccountStatus),
			"createTime":    tea.StringValue(_addr.CreateTime),
			"dailyReqCount": tea.StringValue(_addr.DailyReqCount),
			"monthReqCount": tea.StringValue(_addr.MonthReqCount),
		}
		mas = append(mas, addr)
	}

	return mas, nil
}

//创建发信地址
func (ali *AliClient) CreateMailAddress(accountName string, sendType string) (_err error) {

	createMailAddressRequest := &dm20151123.CreateMailAddressRequest{
		AccountName: tea.String(accountName),
		Sendtype:    tea.String(sendType),
	}
	//执行操作
	_, _err = ali.dmClient.CreateMailAddress(createMailAddressRequest)
	if _err != nil {
		return _err
	}
	return _err
}

//删除指定发信地址
func (ali *AliClient) DeleteMailAddress(emailAddrId string) (_err error) {

	var id int32 = 0
	_id, _err := strconv.Atoi(emailAddrId)
	if _err != nil {
		return _err
	}
	id = int32(_id)
	deleteMailAddressRequest := &dm20151123.DeleteMailAddressRequest{
		MailAddressId: tea.Int32(id),
	}
	//执行操作
	_, _err = ali.dmClient.DeleteMailAddress(deleteMailAddressRequest)
	if _err != nil {
		return _err
	}
	return _err
}

//查询收信人列表
func (ali *AliClient) QueryReceiverByParam() (receives []interface{}, _err error) {

	queryReceiverByParamRequest := &dm20151123.QueryReceiverByParamRequest{}
	//执行操作
	recvs, _err := ali.dmClient.QueryReceiverByParam(queryReceiverByParamRequest)
	if _err != nil {
		return nil, _err
	}

	for _, recv := range recvs.Body.Data.Receiver {
		_recv := map[string]interface{}{
			"id":         tea.StringValue(recv.ReceiverId),
			"name":       tea.StringValue(recv.ReceiversName),
			"status":     tea.StringValue(recv.ReceiversStatus),
			"alias":      tea.StringValue(recv.ReceiversAlias),
			"desc":       tea.StringValue(recv.Desc),
			"createTime": tea.StringValue(recv.CreateTime),
			"count":      tea.StringValue(recv.Count),
		}

		receives = append(receives, _recv)
	}

	return receives, nil
}

//查询收信人明细
func (ali *AliClient) QueryReceiverDetail(recvId string) (users []interface{}, _err error) {

	queryReceiverDetailRequest := &dm20151123.QueryReceiverDetailRequest{
		ReceiverId: tea.String(recvId),
	}
	//执行操作
	us, _err := ali.dmClient.QueryReceiverDetail(queryReceiverDetailRequest)
	if _err != nil {
		return nil, _err
	}
	for _, user := range us.Body.Data.Detail {
		_user := map[string]interface{}{
			"email":      tea.StringValue(user.Email),
			"data":       tea.StringValue(user.Data),
			"createTime": tea.StringValue(user.CreateTime),
		}
		users = append(users, _user)
	}

	return users, nil
}

//创建收信人列表
func (ali *AliClient) CreateReceiver(receiversName, alias string) (_err error) {

	createReceiverRequest := &dm20151123.CreateReceiverRequest{
		ReceiversName:  tea.String(receiversName),
		ReceiversAlias: tea.String(alias),
	}
	//执行操作
	_, _err = ali.dmClient.CreateReceiver(createReceiverRequest)
	if _err != nil {
		return _err
	}
	return _err
}

//删除收信人列表
func (ali *AliClient) DeleteReceiver(receiverId string) (_err error) {

	deleteReceiverRequest := &dm20151123.DeleteReceiverRequest{
		ReceiverId: tea.String(receiverId),
	}
	//执行操作
	_, _err = ali.dmClient.DeleteReceiver(deleteReceiverRequest)
	if _err != nil {
		return _err
	}
	return _err
}

//保存收信人到指定收信人列表
func (ali *AliClient) SaveReceiverDetail(receiverId string, detail string) (_err error) {
	saveReceiverDetailRequest := &dm20151123.SaveReceiverDetailRequest{
		ReceiverId: tea.String(receiverId),
		Detail:     tea.String(detail),
	}
	//执行操作
	_, _err = ali.dmClient.SaveReceiverDetail(saveReceiverDetailRequest)
	if _err != nil {
		return _err
	}
	return _err
}

//从收信人列表删除指定的收信人
func (ali *AliClient) DeleteReceiverDetail(receiverId, email string) (_err error) {

	deleteReceiverDetailRequest := &dm20151123.DeleteReceiverDetailRequest{
		ReceiverId: tea.String(receiverId),
		Email:      tea.String(email),
	}
	//执行操作
	_, _err = ali.dmClient.DeleteReceiverDetail(deleteReceiverDetailRequest)
	if _err != nil {
		return _err
	}
	return _err
}
