package app

import (
	"fmt"
	"noticeservice/global"
	"time"
)

func WriteErrorlog(format string, a ...interface{}) {
	if global.ServerSetting.LogLevel > 0 {
		szLog := fmt.Sprintf(format, a...)
		szInfo := fmt.Sprintf("%s %s\n\n", time.Now().Format("2006-01-02 15:04:05"), szLog)
		if global.ServerSetting.LogMode == 0 {
			fmt.Println(szInfo)
		} else if global.ServerSetting.LogMode == 1 {
			global.Errorlog.Write([]byte(szInfo))
		}
	}
}
func WriteWorkLog(format string, a ...interface{}) {
	if global.ServerSetting.LogLevel > 1 {
		szLog := fmt.Sprintf(format, a...)
		szInfo := fmt.Sprintf("%s %s\n\n", time.Now().Format("2006-01-02 15:04:05"), szLog)
		if global.ServerSetting.LogMode == 0 {
			fmt.Println(szInfo)
		} else if global.ServerSetting.LogMode == 1 {
			global.Worklog.Write([]byte(szInfo))
		}
	}
}
func WriteNetLog(format string, a ...interface{}) {
	if global.ServerSetting.LogLevel > 1 {
		szLog := fmt.Sprintf(format, a...)
		szInfo := fmt.Sprintf("%s %s\n\n", time.Now().Format("2006-01-02 15:04:05"), szLog)
		if global.ServerSetting.LogMode == 0 {
			fmt.Println(szInfo)
		} else if global.ServerSetting.LogMode == 1 {
			global.Netlog.Write([]byte(szInfo))
		}
	}
}
