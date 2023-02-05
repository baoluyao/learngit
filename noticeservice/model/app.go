package model

import "github.com/google/uuid"

type AppHttpRequest struct {
	Todo        string      `json:"td" form:"td" binding:"required"`
	Auth        Auth        `json:"auth" form:"auth" binding:"required"`
	ToId        string      `json:"toId" form:"toId"`
	CmdName     string      `json:"cmdName" form:"cmdName" binding:"required"`
	PayloadType string      `json:"payloadType" form:"payloadType" binding:"required"`
	Payload     interface{} `json:"payload" form:"payload"`
	App         string      `json:"app" form:"app"`
}

type Auth struct {
	With   string `json:"with"`
	UserId string `json:"userid"`
	Token  string `json:"token"`
}

type GeneralAuth struct {
	Auth string `json:"auth" form:"td" binding:"required"`
}

type GetCommonListRequest struct {
	//Filter     map[string]interface{} `json:"filter"`
	Sort *Sort `json:"sort"`
	//MultiSort  []*Sort                `json:"multiSort"`
	Pagination *Pager `json:"page"`
}

type Pager struct {
	CurPage   int `json:"curPage"`         //页码
	PageSize  int `json:"pageSize"`        //每页容量
	TotalRows int `json:"total,omitempty"` //总行数
}

type Sort struct {
	Order string `json:"order"` //排序字段
	Desc  bool   `json:"desc"`  //是否降序
}

type AppNoticeMsg struct {
	MsgId         uuid.UUID `json:"msgId"`
	NoticeCode    string    `json:"noticeCode"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	JumpLink      string    `json:"JumpLink"`
	ReadStatus    bool      `json:"readStatus"`
	NoticeTime    int64     `json:"noticeTime"`
	NoticeTime_tz int64     `json:"noticeTime_tz"`
}

type RobotMsgList struct {
	RobotId     string `json:"robotId"`
	OrgId       string `json:"orgId"`
	UnReadCount int    `json:"unReadCount"`
	TotalCount  int    `json:"totalCount"`
}

type GetNoticeMsg struct {
	RobotId string `json:"robotId"`
	OrgId   string `json:"orgId"`
	Sort    *Sort  `json:"sort"`
	Page    *Pager `json:"page"`
}
