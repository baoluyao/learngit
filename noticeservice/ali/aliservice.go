package ali

var AliService *AliClient

// //发送短信
// func SendSmsTest() {
// 	aliClient := ali.NewClient()

// 	//SMS_241355231   机器故障   ${name}您好!机器人${robot}在${time}有故障产生，请及时排查。
// 	var phoneNums []string = make([]string, 0)
// 	phoneNums = append(phoneNums, "18252102339")
// 	//phoneNums = append(phoneNums, "18252102339") //13771743432

// 	//phoneNums 支持多个号码（是否就是批量发送）   name这里怎么处理，不显示还是用变量
// 	content := map[string]interface{}{
// 		"name":  "11111111112222222222333333333344444444445555555555666666666677777777778888888888999999999900000000001111111111222222222233333333334444444444555555555566666666667777777777888888888899999999990000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999000000000011111111112222222222333333333344444444445555555555666666666677777777778888888888999999999900000000001111111111222222222233333333334444444444555555555566666666667777777777888888888899999999990000000000",
// 		"robot": "1111111111222222222233333333334444444444555555555566666666667777777777888888888899999999990000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999000000000011111111112222222222333333333344444444445555555555666666666677777777778888888888999999999900000000001111111111222222222233333333334444444444555555555566666666667777777777888888888899999999990000000000111111111122222222223333333333444444444455555555556666666666",
// 		"time":  "",
// 		//"time":  time.Now().Format("2006-01-02 15:04:05"),
// 	}
// 	rawByte, err := json.Marshal(content)
// 	if err != nil {
// 		fmt.Println("error:", err.Error())
// 	} else {
// 		if code, msg, err := aliClient.SendSms(phoneNums, "科沃斯商用机器人", "SMS_241355231", string(rawByte)); err != nil {
// 			fmt.Println("error:", err.Error())
// 		} else {
// 			if code != "" {
// 				fmt.Println("code:", code, "msg:", msg)
// 			}
// 		}
// 	}

// }

// //批量发送短信      模板相同  模板中的参数变量不同 来达到发送不同短信内容的目的
// func SendBatchSmsTest() {
// 	//aliClient := ali.NewClient()  aliClient被声明为全局变量  不用在方法中一个一个声明

// 	//在批量方法中测试发送单封邮件  （支持）
// 	// phoneNumberJson := `["15905101081"]`
// 	// signNameJson := `["科沃斯商用机器人"]`
// 	// templateCode := `SMS_241355231`
// 	// templateParamJson := `[{"name":"鲍路尧","robot":"1002","time":""}]`

// 	phoneNumberJson := `["15905101081","18252102339"]`
// 	signNameJson := `["科沃斯商用机器人","科沃斯商用机器人"]`
// 	templateCode := `SMS_241355231`
// 	templateParamJson := `[{"name":"鲍路尧","robot":"1002","time":""},{"name":"周红艳","robot":"1001","time":""}]`

// 	code, msg, err := aliClient.SendBatchSms(phoneNumberJson, signNameJson, templateCode, templateParamJson)
// 	if err != nil {
// 		fmt.Println("error:", err.Error())
// 	}
// 	fmt.Println("code:", code, "msg:", msg)
// }

// //单封邮件          后面需不要建一个回信地址  addresstype 取0和1 没区别
// func SingleSendMail() {
// 	//f700b8c82c097537cbc6a578dbce3b4c  receiverid
// 	accountName := `public@marketing.ecovacs-c.com`
// 	//addressType:=1
// 	toAddress := `luyao.bao@ecovacs.com`
// 	subject := `异常提醒通知`
// 	//replyToAddress:true
// 	textBody := `尊敬的鲍路尧先生您好，您维护的机器人19980008发生异常，请及时处理fdafadsfadsfadfsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssfdafadsfadsfadfsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssfdafadsfadsfadfsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssfdafadsfadsfadfsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssfdafadsfadsfadfsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssfdafadsfadsfadfsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssfdafadsfadsfadfsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss`
// 	err := aliClient.SingleSendMail(textBody, toAddress, subject, accountName)
// 	if err != nil {
// 		fmt.Println("发送单封邮件发生异常，error:%v", err)
// 		return
// 	}
// 	fmt.Println("邮件发送成功，请及时查看！")
// }

// //批量发送邮件
// func SendBatchMail() {
// 	templateName := `ecovscsBiz` //ecovacsBiz
// 	accountName := `public@marketing.ecovacs-c.com`
// 	receiversName := `eco_bly`
// 	err := aliClient.BatchSendMail(templateName, accountName, receiversName)
// 	if err != nil {
// 		fmt.Println("批量发送邮件异常，error:%v", err)
// 		return
// 	}
// 	fmt.Println("批量发送邮件成功，请及时查看！")
// }

// //语音通知
// func SingleCallByTts() {

// 	aliClient := ali.NewClient()
// 	CalledShowNumber := "057128013952"
// 	CalledNumber := "15905101081"
// 	//CalledNumber := "13771743432" //bin
// 	//TtsCode := "TTS_242686821" //文本转语音模板   公共外呼  参数:name|robotsn
// 	//TtsCode := "TTS_242698335" //文本转语音模板   专属外呼  参数：name、robotsn
// 	TtsCode := "TTS_242688526" //文本转语音模板   专属外呼(专属模式外呼2)     ${name}您好，机器人${detail}，感谢您的接听。
// 	//TtsCode := "TTS_242706659"  //语音验证码模板  实际内容仅支持数字和字母
// 	PlayTimes := 3
// 	Volume := 100
// 	Speed := -150

// 	param := map[string]interface{}{}
// 	param["name"] = "王彬先生"
// 	param["detail"] = "此内容为detail详细信息，让我先唱首歌，啦啦啦，啦啦啦，啦啦啦，机器人1998008清洁模块发生异常，请及时处理，此通知一共播报三次。啦啦啦，啦啦啦。邮件 html 正文，限制28K邮件 html 正文，限制28K邮件 html 正文，限制28K邮件 html 正文，限制28K邮件 html 正文，限制28K邮件 html 正文，限制28K邮件 html 正文，限制28K"
// 	// param := map[string]interface{}{
// 	// 	"product": "阿里云登录",
// 	// 	"code":    "123456",
// 	// }
// 	dataString := helper.MapToJson(param)
// 	fmt.Printf(string(dataString))
// 	if code, msg, err := aliClient.SingleCallByTts(CalledShowNumber, CalledNumber, TtsCode,
// 		string(dataString), PlayTimes, Volume, Speed); err != nil {
// 		fmt.Println("error:", err.Error())
// 	} else {
// 		if code != "" {
// 			fmt.Println("code:", code, "msg:", msg)
// 		}
// 	}
// }

// //ssdb存取
// func SSDBTest() {
// 	//step1： 测试读取cockroach数据库数据
// 	d := dao.New(global.DBEngine)
// 	var list []*model.NoticeHead
// 	args := []interface{}{}
// 	list, err := d.GetNoticeHeadList(args)
// 	if err != nil {
// 		return
// 	}
// 	var s string
// 	var marshal []byte
// 	if len(list) > 0 {
// 		marshal, err = json.Marshal(*list[0])
// 		if err != nil {
// 			return
// 		}
// 		fmt.Println(string(marshal))
// 		s = string(marshal)
// 	}
// 	//step2:  数据存入ssdb
// 	curTime := time.Now().Format("2006-01-02 15:04:05") //hash
// 	hashKey := "noticedata:noticehead"
// 	_, err1 := cachelib.HSET(hashKey, s, curTime)
// 	if err1 != nil {
// 		fmt.Println("设置缓存成功")
// 	}
// 	// key := "kv_noticedata:noticehead" //key-value
// 	// cachelib.SetCommonOnlineInfo(key, *list[0])

// 	key1 := "common:kv_noticedata:noticehead"
// 	//s1, err := cachelib.SET(key1, s)
// 	s1, err := cachelib.SET(key1, "test_update_data")
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(s1)
// 	//step3:  读取数据
// 	readKey := "kv_noticedata:noticehead"
// 	byteData, err2 := cachelib.GetCommonOnlineInfo(readKey)
// 	if err2 != nil {
// 		fmt.Println("GetCommonOnlineInfo error:%v", err2)
// 		return
// 	}
// 	fmt.Println(string(byteData))
// }
