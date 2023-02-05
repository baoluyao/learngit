package routers

import (
	"encoding/json"
	"fmt"
	"noticeservice/helper/service"
	"noticeservice/model"
	"noticeservice/pkg/app"

	"github.com/carr123/fmx"
)

func GetNoticeMsg(c *fmx.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	datas := c.ReadBody()
	fmt.Println(datas)
	// var param model.AppHttpRequest

	// if err := json.Unmarshal(data, &param); err != nil {
	// 	app.WriteErrorlog("GetNoticeMsg bodyjson err:%v", err)
	// }

	var param model.GetNoticeMsg

	data := c.MustGet("payload")
	fmt.Println(data)
	if err := json.Unmarshal(data.([]byte), &param); err != nil {
		app.WriteErrorlog("GetNoticeMsg bodyjson err:%v", err)
	}

	records, page, err := svc.GetNoticeMsg(&param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(fmx.H{"msgs": records, "page": page})
}

func ReadMsg(c *fmx.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	var payload struct {
		RobotId  string `json:"robotId"`
		OrgId    string `json:"orgId"`
		ReadType int    `json:"readType"`
		MsgId    string `json:"msgId",omitempty`
	}

	data := c.MustGet("payload")
	if err := json.Unmarshal(data.([]byte), &payload); err != nil {
		app.WriteErrorlog("GetNoticeMsg bodyjson err:%v", err)
	}
	if payload.ReadType == 0 {
		err := svc.ReadSingleMsg(payload.MsgId)
		if err != nil {
			response.ToErrorResponse(err)
		}
	} else {
		err := svc.ReadAllMsg(payload.RobotId)
		if err != nil {
			response.ToErrorResponse(err)
		}
	}

	response.ToResponse()
}

func GetUserRobotsMsg(c *fmx.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	data := c.ReadBody()
	var param model.AppHttpRequest

	if err := json.Unmarshal(data, &param); err != nil {
		app.WriteErrorlog("GetNoticeMsg bodyjson err:%v", err)
	}

	//userid -> did
	did, err := svc.GetBindedDevices(param.Auth.UserId)
	if err != nil {
		response.ToErrorResponse(err)
	}
	//E0AD12345D0000000084
	//dids -> robotsn/robotid
	robotId, err := svc.GetRobotSNByDid(did)
	if err != nil {
		response.ToErrorResponse(err)
	}

	//robotsn -> msgs
	robotlist, err := svc.GetRobotsMsg(robotId)
	if err != nil {
		response.ToErrorResponse(err)
	}

	response.ToResponse(robotlist)
}
