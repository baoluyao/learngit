package service

import (
	"encoding/json"
	"fmt"
	"noticeservice/ali"
	"noticeservice/global"
	"noticeservice/helper/dao"
	"noticeservice/helper/tool"
	"noticeservice/model"
	"noticeservice/pkg/app"
	"noticeservice/pkg/errcode"
	"noticeservice/pkg/helper"
	"sync"
	"time"

	"github.com/carr123/fmx"
)

var (
	aliClient ali.AliClient //声明一个全局变量？？  哪个方法里面都可以用
)

func SendSmsTest(c *fmx.Context) {
	var startTime, finishTime string
	startTime = time.Now().Format(global.TimeLayout12)

	response := app.NewResponse(c)
	//svc := service.New(c.Request.Context())
	var param struct {
		RobotSN     string `json:"robotSN"`
		NoticeCode  string `json:"noticeCode"`
		ProductCode string `json:"productCode"`
		OrgId       string `json:"orgId"`
	}
	data := c.ReadBody()
	if err := json.Unmarshal(data, &param); err != nil {
		app.WriteErrorlog("SendSmsTest bodyjson err:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	var wg sync.WaitGroup
	//aliClient := ali.NewClient()
	dao := dao.New(global.DBEngine)
	contendIds, err := dao.GetContentIdsByBobotInfo(param.OrgId, param.ProductCode)
	if err != nil {
		app.WriteErrorlog("dao.GetContentIdsByBobotInfo error:%v", err)
	}

	if len(contendIds) > 0 {

		fmt.Println(len(contendIds))
		wg.Add(len(contendIds))

		for _, contentId := range contendIds {
			_contentId := contentId //解决go中的经典问题
			//fmt.Println(_contentId)
			// fmt.Println(&_contentId)
			go func() {
				defer wg.Done()
				//通知谁   通知方式
				scheme, err := dao.GetSchemeDetailByContentId(_contentId) //默认一个contentId只能对应一个scheme
				if err != nil {
					app.WriteErrorlog("dao.GetSchemeDetailByContentId error:%v", err)
				}

				//1.通知谁
				users, err := dao.GetUsersByRoleIds(scheme.RoleIds, param.OrgId)
				if err != nil {
					app.WriteErrorlog("dao GetUsersByRoleIds error:%v", err)

				}
				//2.通知方式  那个code数组中有这个 noticeCode
				channels, err := dao.GetNoticeChannels(param.NoticeCode, _contentId)
				if err != nil {
					app.WriteErrorlog("dao GetNoticeChannels error:%v", err)
				}

				fmt.Println(_contentId.String(), channels.StationMsg.String(), channels.Msg.String(), channels.Phone.String())
				//通知内容
				contentDetail, err := dao.GetContentDetail(param.NoticeCode, param.ProductCode, _contentId)
				if err != nil {
					app.WriteErrorlog("dao.GetContentDetail error%v", err)
				}

				// users contentDetail  channels  三者必须都不为nil   才满足发通知的条件
				if users == nil {
					app.WriteErrorlog("users nil !  "+_contentId.String()+"  schemeid:%v,roleids:%v", scheme.SchemeId, scheme.RoleIds)
				} else if contentDetail == nil {
					app.WriteErrorlog("contentdetail nil!  "+_contentId.String()+" schemeid:%v,contentid:%v,productcode:%v", scheme.SchemeId, _contentId, param.ProductCode)
				} else if channels == nil {
					app.WriteErrorlog("channels nil !  "+_contentId.String()+"  schemeid:%v,noticecode:%v", scheme.SchemeId, param.NoticeCode)
				} else {
					HandleNoticeEvents(channels, users, contentDetail)
				}

			}()

		}
	}

	wg.Wait()
	//time.Sleep(3 * time.Second)
	finishTime = time.Now().Format(global.TimeLayout12)
	lastTime := tool.StringToTimestamp(global.TimeLayout12, finishTime) - tool.StringToTimestamp(global.TimeLayout12, startTime)
	fmt.Println("notice success")
	c.JSON(200, fmx.H{"res": "success", "startTime:": startTime, "finishTime": finishTime, "lastTime": lastTime})
}

func HandleNoticeEvents(channel *model.NoticeChannel, users []*model.Notice_Users, contentDetail *model.Notice_ContentDetail) {
	if channel.StationMsg.String() != "" {

	}
	if channel.Msg.String() != "" {
		SendSms(users, contentDetail)
		//go SendSms(users, contentDetail)
	}
	if channel.Phone.String() != "" {
		for _, user := range users {
			_user := user //解决go中的经典问题
			SingleCallByTts(_user, contentDetail)
			//go SingleCallByTts(_user, contentDetail)
		}
	}
}

func SendSms(users []*model.Notice_Users, contentDetail *model.Notice_ContentDetail) {
	var phoneNumber, signName []model.STRING
	var templateParam []map[string]interface{}
	for _, user := range users {
		phoneNumber = append(phoneNumber, user.Phone)
		signName = append(signName, model.SetString("科沃斯商用机器人"))
		templateParam = append(templateParam, map[string]interface{}{
			"name":  user.UserName,
			"robot": "19980008" + contentDetail.Title.String() + contentDetail.Content.String() + contentDetail.Solution.String() + contentDetail.JumpLink.String(),
			"time":  time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	phoneNumberBytes, err := json.Marshal(phoneNumber)
	if err != nil {
		app.WriteErrorlog("")
	}
	signNameBytes, err := json.Marshal(signName)
	if err != nil {
		app.WriteErrorlog("")
	}
	templateParamBytes, err := json.Marshal(templateParam)
	if err != nil {
		app.WriteErrorlog("")
	}
	phoneNumberJson := string(phoneNumberBytes)
	signNameJson := string(signNameBytes)
	templateParamJson := string(templateParamBytes)

	//aliClient.SendBatchSms(phoneNumberJson, signNameJson, templateCode, templateParamJson)
	//signName := "科沃斯商用机器人"
	templateCode := "SMS_241355231"

	code, msg, err := ali.AliService.SendBatchSms(phoneNumberJson, signNameJson, templateCode, templateParamJson)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println("code:", code, "msg:", msg)
}

func SingleCallByTts(user *model.Notice_Users, contentDetail *model.Notice_ContentDetail) {

	CalledShowNumber := "057128013952"
	CalledNumber := user.Phone.String()
	//CalledNumber := "13771743432" //bin
	//TtsCode := "TTS_242686821" //文本转语音模板   公共外呼  参数:name|robotsn
	//TtsCode := "TTS_242698335" //文本转语音模板   专属外呼  参数：name、robotsn
	TtsCode := "TTS_242688526" //文本转语音模板   专属外呼(专属模式外呼2)     ${name}您好，机器人${detail}，感谢您的接听。
	//TtsCode := "TTS_242706659"  //语音验证码模板  实际内容仅支持数字和字母
	PlayTimes := 2
	Volume := 100
	Speed := -50

	param := map[string]interface{}{}
	param["name"] = user.UserName
	param["detail"] = contentDetail.Title.String() + contentDetail.Content.String() + contentDetail.Solution.String() + contentDetail.JumpLink.String()
	// param := map[string]interface{}{
	// 	"product": "阿里云登录",
	// 	"code":    "123456",
	// }
	dataString := helper.MapToJson(param)
	fmt.Printf(string(dataString))
	if code, msg, err := ali.AliService.SingleCallByTts(CalledShowNumber, CalledNumber, TtsCode,
		string(dataString), PlayTimes, Volume, Speed); err != nil {
		fmt.Println("error:", err.Error())
	} else {
		if code != "" {
			fmt.Println("code:", code, "msg:", msg)
		}
	}
}

//短信通知 和 语音通知分别开两个协程

//单条语音通知之间分别开不同的协程 （多个可换成创建语音通话任务 ）

//ps： 非main函数结束，一子协程不会挂掉  会继续执行
