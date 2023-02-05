package service

import (
	"strings"
	"time"

	"noticeservice/global"
	"noticeservice/pkg/app"
	"noticeservice/pkg/batchwriter"

	"github.com/nats-io/nats.go"
)

var bw *batchwriter.AsyncBatchWriter

type Topic struct {
	Topic    string                  `json:"topic"`
	Function func(bin []byte) []byte `json:"function"`
}

type topicinfo struct {
	Topic    string                  `json:"topic"`
	Function func(bin []byte) []byte `json:"function"`
}

//订阅机器上报rn  statusOption

func InitBroadcast() error {
	var topics []topicinfo
	topics = append(topics,
		topicinfo{
			Topic:    strings.ToUpper("BROADCAST_eventRep.*"), //订阅获取产品下机器列表
			Function: HandlerEventRep,
		})

	for _, subtopic := range topics {
		if _, err := global.Nats_client.QueueSubscribe(subtopic.Topic, "SOCSERVER", RPC_HANDLER(subtopic.Function)); err != nil {
			app.WriteErrorlog("service.InitBroadcast.%s error %v", subtopic, err)
			return err
		}
	}
	return nil
}

func RPC_HANDLER(handler func(bin []byte) []byte) func(msg *nats.Msg) {
	pool := app.NewWorkerPool(18)
	pool.SetMaxGoroutine(256)
	pool.SetMaxIdleGoroutine(0)
	pool.SetMaxIdleTime(time.Second * 60)
	pool.SetHandler(func(a interface{}) {
		defer global.Recover()
		msg := a.(*nats.Msg)
		app.WriteWorkLog("RPC_HANDLER %s bodyjson:%v", msg.Subject, string(msg.Data))
		res := handler(msg.Data)
		if res != nil {
			msg.Respond(res)
		}
	})
	return func(msg *nats.Msg) {

		pool.PushItem(msg)
	}
}

//处理数据  1.先通知   2.后存库
func HandlerEventRep(bin []byte) []byte {
	// var param model.Msgdata
	// if err := json.Unmarshal(bin, &param); err != nil {
	// 	app.WriteErrorlog("UpdateRobotBattary %v bodyjson err:%v", err)
	// 	return nil
	// }
	// fmt.Println(param)
	// EventData := &model.EventData{}

	// err := app.MapToStruct(param.Data, EventData)
	// if err != nil {
	// 	app.WriteErrorlog("BaseStationReportController error,err:%v", err.Error())
	// 	return nil
	// }
	// fmt.Println(EventData)
	// //URORA test 极光推送
	// appMsgPusher := jpush.NewJPush()
	// fmt.Println(EventData.Events)
	// for _, v := range EventData.Events {
	// 	appMsgPusher.JpushDemo(v)
	// }

	return nil
}
