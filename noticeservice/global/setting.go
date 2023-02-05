package global

import (
	easylog "noticeservice/pkg/logger"
	"noticeservice/pkg/setting"
)

var (
	DirPath string

	ServerSetting *setting.ServerSettings
	AliSetting    *setting.AliSettings

	NgiotSetting    *setting.NgiotSettings
	DatabaseSetting *setting.DatabaseSettings

	NatsSetting *setting.NatsSettings
	SSDBSetting *setting.SSDBSettings

	Errorlog *easylog.EasyLog
	Worklog  *easylog.EasyLog
	Netlog   *easylog.EasyLog
)
