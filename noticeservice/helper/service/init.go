package service

import (
	"noticeservice/global"
	"noticeservice/pkg/app"
	"noticeservice/pkg/setting"
	"time"

	easylog "noticeservice/pkg/logger"

	dbserver "github.com/carr123/easysql/cockroach"
	"github.com/nats-io/nats.go"
)

//读取配置文件
func SetupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Ali", &global.AliSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("ngiot", &global.NgiotSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("nats", &global.NatsSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("SSDB", &global.SSDBSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func Initlog() {
	global.Errorlog = easylog.NewLog(1000, time.Second)
	global.Errorlog.SetDir(global.ServerSetting.LogSavePath, "errorlog.txt")
	global.Errorlog.SetMaxFileSize(global.ServerSetting.LogFileSize) //1MB for every file
	global.Errorlog.SetMaxFileCount(global.ServerSetting.LogFileCount)

	global.Worklog = easylog.NewLog(1000, time.Second)
	global.Worklog.SetDir(global.ServerSetting.LogSavePath, "worklog.txt")
	global.Worklog.SetMaxFileSize(global.ServerSetting.LogFileSize * 1024 * 1024) //1MB for every file
	global.Worklog.SetMaxFileCount(global.ServerSetting.LogFileCount)

	global.Netlog = easylog.NewLog(1000, time.Second)
	global.Netlog.SetDir(global.ServerSetting.LogSavePath, "netlog.txt")
	global.Netlog.SetMaxFileSize(global.ServerSetting.LogFileSize) //1MB for every file
	global.Netlog.SetMaxFileCount(global.ServerSetting.LogFileCount)
}

//初始化数据库连接
func SetupDBEngine() error {
	var err error
	global.DBEngine, err = dbserver.New(global.DatabaseSetting.Addr, global.DatabaseSetting.MaxIdleConns)
	if err != nil {
		app.WriteErrorlog("SetupDBEngine err:sql:%s", err, global.DatabaseSetting.Addr)
		return err
	}
	return nil
}

func Init_nates() error {
	var err error
	global.Nats_client, err = nats.Connect(
		global.NatsSetting.Connect,
		nats.Name(global.NatsSetting.Name), //连接名称
		nats.Timeout(global.NatsSetting.Timeout*time.Second),
		nats.PingInterval(global.NatsSetting.PingInterval*time.Second),
		nats.MaxPingsOutstanding(global.NatsSetting.MaxPingsOutstanding),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(global.NatsSetting.ReconnectWait*time.Second),
		nats.ReconnectBufSize(global.NatsSetting.ReconnectBufSize*1024*1024), // Set reconnect buffer size in bytes (5 MB), 服务器不在线时客户端发送缓冲区的大小
		nats.UserInfo(global.NatsSetting.UserName, global.NatsSetting.Password),
	)
	if err != nil {
		return err
	}
	// global.JSM, err = global.Nats_client.JetStream()
	// if err != nil {
	// 	return err
	// }

	err = InitBroadcast()
	if err != nil {
		return err
	}
	return nil
}
