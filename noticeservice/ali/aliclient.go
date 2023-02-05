package ali

import (
	"noticeservice/global"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	dyvmsapi20170525 "github.com/alibabacloud-go/dyvmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

//
type AliClient struct {
	dmClient    *dm20151123.Client       //邮件
	dysmsClient *dysmsapi20170525.Client //短信
	voiceClient *dyvmsapi20170525.Client //语音
}

//每调用一个NewClient就会初始化三个客户端   会不会造成资源浪费   有没有过期时间（即需不需要重新创建） 会不会断连 断连的话就需要一个重连的机制
func NewClient() *AliClient {
	dmClient, err1 := CreateDMClient(tea.String(global.AliSetting.AccessKeyId), tea.String(global.AliSetting.AccessKeySecret))
	dysmsClient, err2 := CreateDYSMSClient(tea.String(global.AliSetting.AccessKeyId), tea.String(global.AliSetting.AccessKeySecret))
	voiceClient, err3 := CreateVoiceClient(tea.String(global.AliSetting.AccessKeyIdVoice), tea.String(global.AliSetting.AccessKeySecretVoice))
	if err1 != nil && err2 != nil && err3 != nil {
		return &AliClient{}
	} else {
		return &AliClient{dmClient: dmClient, dysmsClient: dysmsClient, voiceClient: voiceClient}
	}
}

//创建短信客户端
func CreateDYSMSClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(global.AliSetting.EndpointDysms) //此域名需可配
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

//创建邮件客户端
func CreateDMClient(accessKeyId *string, accessKeySecret *string) (_result *dm20151123.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String(global.AliSetting.EndpointDm)
	_result = &dm20151123.Client{}
	_result, _err = dm20151123.NewClient(config)
	return _result, _err
}

//创建语音客户端
func CreateVoiceClient(accessKeyId *string, accessKeySecret *string) (_result *dyvmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String(global.AliSetting.EndpointVoice)
	_result = &dyvmsapi20170525.Client{}
	_result, _err = dyvmsapi20170525.NewClient(config)
	return _result, _err
}
