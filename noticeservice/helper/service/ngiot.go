package service

import (
	"encoding/json"
	"noticeservice/global"
	"noticeservice/pkg/app"
	"noticeservice/pkg/errcode"

	"github.com/carr123/fmx"
)

func (svc *Service) GetBindedDevices(userid string) ([]string, *errcode.Error) {
	param := fmx.H{
		"auth": fmx.H{
			"with":        "oauth",
			"accesstoken": global.SocUser_Token,
		},
		"td":   "GetBindedDevices",
		"user": userid,
	}
	resdata, statusCode, err := app.HttpMethod("POST",
		global.NgiotSetting.Addr+"/api/dim/devmanager.do", param)
	if err != nil || statusCode != 200 {
		app.WriteErrorlog("svc.GetBindedDevices params:%v err:%v,statuscode:%d", param, err, statusCode)
		return nil, errcode.NgIOTNetError

	}
	var rejson map[string]interface{}
	if err := json.Unmarshal(resdata, &rejson); err != nil {
		app.WriteErrorlog("svc.GetBindedDevices Unmarshal params:%v data  err:%s,", param, string(resdata))
		return nil, errcode.NgiotGetBindDevicesError
	}
	app.WriteWorkLog("svc.GetBindedDevices params:%v get response:%s", param, string(resdata))
	ret, ok := rejson["ret"]
	if ok && ret.(string) == "ok" {
		var did []string
		devices := rejson["devices"].([]interface{})
		if len(devices) > 0 {
			for _, device := range devices {
				did = append(did, device.(map[string]interface{})["id"].(string))
				//resources = append(resources, device.(map[string]interface{})["resource"].(string))
			}
		}
		return did, nil
	} else {
		app.WriteErrorlog("svc.GetBindedDevices param:%v error err:%s,", param, string(resdata))
		return nil, errcode.NgiotGetBindDevicesError //rejson["error"].(string)
	}
}
