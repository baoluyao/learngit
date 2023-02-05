package model

import "github.com/google/uuid"

type Notice_Schemes struct {
	SchemeId        uuid.UUID   `json:"schemeId"`
	OrgId           uuid.UUID   `json:"orgId"`
	ContentId       uuid.UUID   `json:"contentId"`
	RoleIds         StringArray `json:"roleIds" db:"roleids"`
	UserIds         StringArray `json:"userIds" db:"userids"`
	StationMsgCodes StringArray `json:"staionMsgCodes"`
	MsgCodes        StringArray `json:"msgCodes"`
	PhoneCodes      StringArray `json:"phoneCodes"`
	*Model
}

//通知内容方案详细
type Notice_ContentHead struct {
	ContentId   uuid.UUID   `json:"contendId"`
	ProductCode StringArray `json:"productCode"`
	ContentName STRING      `json:"contentName"`
}

//通知内容方案详细
type Notice_ContentDetail struct {
	ContentId   uuid.UUID `json:"contentId"`
	ProductCode STRING    `json:"productCode"`
	NoticeCode  STRING    `json:"noticeCode"`
	Title       STRING    `json:"title"`
	Content     STRING    `json:"content"`
	Solution    STRING    `json:"solution"`
	JumpLink    STRING    `json:"jumpLink"`
}

//user
type Notice_Users struct {
	UserId   uuid.UUID `json:"userId" db:"userid"`
	UserName STRING    `json:"userName" db:"username"`
	Email    STRING    `json:"email" db:"email"`
	Phone    STRING    `json:"phone" db:"phone"`
}

type NoticeChannel struct {
	StationMsg STRING `json:"staionMsg"`
	Msg        STRING `json:"msg"`
	Phone      STRING `json:"phone"`
}
