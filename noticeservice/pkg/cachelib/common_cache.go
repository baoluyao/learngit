package cachelib

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	commonOnlinePrefix = "common:"
)

func getCommonOnlineKey(k string) (key string) {
	key = fmt.Sprintf("%s%s", commonOnlinePrefix, k)

	return
}

// 获取缓存数据
func GetCommonOnlineInfo(k string) (e []byte, err error) {

	key := getCommonOnlineKey(k)
	data, err := GET(key)
	if err != nil {
		if err.Error() == Nil {
			logrus.Infoln("获取在线数据 GetCommonOnlineInfo", key, err.Error())

			return
		}

		logrus.Infoln("获取在线数据 GetCommonOnlineInfo", key, err.Error())

		return
	}

	e = []byte(data)

	return
}

// 设置缓存数据
func SetCommonOnlineInfo(k string, param interface{}) (err error) {

	key := getCommonOnlineKey(k)
    //json.Unmarshal()
	valueByte, err := json.Marshal(param)
	if err != nil {
		logrus.Infoln("设置在线数据 SetCommonOnlineInfo json Marshal", key, err.Error())

		return
	}

	_, err = SET(key, string(valueByte)) //将value关联到key
	if err != nil {
		logrus.Infoln("设置在线数据 SetCommonOnlineInfo 失败", key, err.Error())

		return
	}

	return
}

// 删除缓存数据
func DelCommonOnlineInfo(k string) (bool, error) {
	key := getCommonOnlineKey(k)
	return DEL(key)
}
