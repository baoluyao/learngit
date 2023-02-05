package tool

import (
	"math/rand"
	"time"
)

// StringToTimestamp 字符串转时间戳
func StringToTimestamp(timeLayout string, timeString string) int64 {
	dateTime, err := time.Parse(timeLayout, timeString)
	if err != nil {
		return 0
	}
	timestamp := dateTime.Unix()
	return timestamp
}

// type Soctime struct {
// 	time.Time
// }

// func (t Soctime) GetTimeStamp() int64 {
// 	return t.UnixMilli()
// }

// TimestampToString 时间戳转字符串
func TimestampToString(timeLayout string, timestamp int64) string {
	dateTime := time.Unix(timestamp, 0).Format(timeLayout)
	return dateTime
}
func GetRandString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//modify by zouchenchen
// func SetPubModule(modulecode string) string {
// 	// modulegroup := global.ModuleSetting
// 	for key, v := range *global.ModuleSetting {
// 		modulegroup := strings.Split(v.(string), ",")
// 		for _, v := range modulegroup {
// 			if v == modulecode {
// 				return key
// 			}
// 		}
// 	}
// 	return modulecode
// }

// func SetPubModule(modulecode string) string {
// 	// if modulecode == "x86" || modulecode == "apk" {
// 	// 	return modulecode
// 	// }
// 	if global.ModuleSetting.X86 != "" {
// 		x86arr := strings.Split(global.ModuleSetting.X86, ",")
// 		for _, v := range x86arr {
// 			if v == modulecode {
// 				return "x86"
// 			}
// 		}
// 	}
// 	if global.ModuleSetting.Apk != "" {
// 		apkarr := strings.Split(global.ModuleSetting.Apk, ",")
// 		for _, v := range apkarr {
// 			if v == modulecode {
// 				return "apk"
// 			}
// 		}
// 	}
// 	return modulecode
// }
