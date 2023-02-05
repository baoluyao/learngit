package service

import (
	"noticeservice/model"
	"noticeservice/pkg/errcode"
)

//region  ==========app通知中心==========

func (svc *Service) GetNoticeMsg(param *model.GetNoticeMsg) ([]*model.AppNoticeMsg, *model.Pager, *errcode.Error) {
	return svc.dao.GetNoticeMsg(param)
}

func (svc *Service) GetRobotSNByDid(did []string) ([]string, *errcode.Error) {
	return svc.dao.GetRobotSNByDid(did)
}

//endregion

//region  ==========app通知中心==========

func (svc *Service) GetRobotsMsg(robotId []string) (*[]model.RobotMsgList, *errcode.Error) {
	return svc.dao.GetRobotsMsg(robotId)
}

func (svc *Service) ReadSingleMsg(msgId string) *errcode.Error {
	return svc.dao.ReadSingleMsg(msgId)
}

func (svc *Service) ReadAllMsg(robotId string) *errcode.Error {
	return svc.dao.ReadAllMsg(robotId)
}

//endregion
