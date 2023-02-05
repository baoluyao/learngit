package ali

import (
	"noticeservice/pkg/app"

	dyvmsapi20170525 "github.com/alibabacloud-go/dyvmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	//util  "github.com/alibabacloud-go/tea-utils/service"
	// openapi "github.com/alibabacloud-go/darabonba-openapi/client"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */

// 语音通知（tts模板）
func (ali *AliClient) SingleCallByTts(callShowNumber, callNumber, ttsCode, ttsParam string, palyTimes int, volume int, speed int) (code, msg string, _err error) {
	// client, _err := CreateVoiceClient(tea.String(AccessKeyId), tea.String(AccessKeySecret))
	// if _err != nil {
	// 	return "", "", _err
	// }

	singleCallByTtsRequest := &dyvmsapi20170525.SingleCallByTtsRequest{
		CalledShowNumber: tea.String(callShowNumber),  //被叫显号
		CalledNumber:     tea.String(callNumber),      //接受语音通知的手机号码
		TtsCode:          tea.String(ttsCode),         //已通过审核的语音验证码模板ID
		TtsParam:         tea.String(ttsParam),        //模板中参数变量     目前只支持两个变量
		PlayTimes:        tea.Int32(int32(palyTimes)), //一通语音电话内容的播放次数 取值范围： 1 ~3 默认取值 3
		Volume:           tea.Int32(int32(volume)),    //音量： 0~100  默认100
		Speed:            tea.Int32(int32(speed)),     //语速： -500~500
		// OutId:            tea.String(outId)        //回执ID将此ID带回给调用方
	}
	// 执行操作
	result, _err := ali.voiceClient.SingleCallByTts(singleCallByTtsRequest)
	if _err != nil {
		app.WriteWorkLog("SingleCallByTts fail, error:%v", _err)
		return "", "", _err
	}

	app.WriteWorkLog("SingleCallByTts success, params:%v", singleCallByTtsRequest)

	code = tea.StringValue(result.Body.Code)
	msg = tea.StringValue(result.Body.Message)

	return code, msg, nil
}

//创建通话任务  （语音通知和语音验证码的号码资源可以支持 一个号码 10个并发。）
// func(ali *AliClient) CreateCallTask(taskName,bizType,templateCode,templateName,resourceType,dataType,resource,data string) (code string ,taskId int,_err error){
// 	createCallTaskRequest:=&dyvmsapi20170525.CreateCallTaskRequest{
// 		TaskName: tea.String(taskName),
// 		BizType: tea.String(bizType),
// 		TemplateCode: tea.String(templateCode),
// 		TemplateName: tea.String(templateName),
// 		ResourceType: tea.String(resourceType),
// 		DataType: tea.String(dataType),
// 		Resource: tea.String(resource),
// 		Data: tea.String(data),
// 	}
// 	runtime:=&util.RuntimeOptions{}
// 	//执行操作
// 	result,_err:=ali.voiceClient.CreateCallTaskWithOptions(createCallTaskRequest,runtime)
// 	if _err != nil {
// 		return "", 0, _err
// 	}
// 	code = tea.StringValue(result.Body.Code)
// 	taskId = tea.Int(result.Body.Data)
//     //go
// 	return code,taskId,nil
// }

//执行通话任务
func (ali *AliClient) ExecuteCallTask(taskId int) {

}

// 语音通知（录音文件）
// func (ali *AliClient) SingleCallByVoice(callShowNumber, callNumber, voiceCode string, playTimes int, volume int, speed int) (code, msg string, _err error) {
// 	// client, _err := CreateVoiceClient(tea.String(AccessKeyId), tea.String(AccessKeySecret))
// 	// if _err != nil {
// 	// 	return "", "", _err
// 	// }

// 	singleCallByVoiceRequest := &dyvmsapi20170525.SingleCallByVoiceRequest{
// 		CalledShowNumber: tea.String(callShowNumber),
// 		CalledNumber:     tea.String(callNumber),
// 		VoiceCode:        tea.String(voiceCode),
// 		PlayTimes:        tea.Int32(int32(playTimes)),
// 		Volume:           tea.Int32(int32(volume)),
// 		Speed:            tea.Int32(int32(speed)),
// 	}
// 	// 执行操作
// 	result, _err := ali.voiceClient.SingleCallByVoice(singleCallByVoiceRequest)
// 	if _err != nil {
// 		return "", "", _err
// 	}

// 	code = tea.StringValue(result.Body.Code)
// 	msg = tea.StringValue(result.Body.Message)

// 	return code, msg, nil
// }
