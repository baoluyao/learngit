package main

import (
	"fmt"
	"net/http"
	"noticeservice/ali"
	"noticeservice/global"
	"noticeservice/helper/service"
	"noticeservice/jpush"
	"noticeservice/pkg/app"
	"noticeservice/pkg/cachelib"
	"noticeservice/routers"
	"time"

	"github.com/carr123/fmx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// var (
// 	aliClient *ali.AliClient //声明一个全局变量？？  哪个方法里面都可以用
// )

var appMsgPusher *jpush.JPUSH //APP端极光推送
// 软件版本号  通知服务
const (
	SW_Version  = "CS-NOTV001.01.1"
	robotSN     = "1998008"
	noticeCode  = "1001"
	productCode = "seaTide"
	orgId       = "2c3f38ff-4436-4a84-8cf0-2d381b350ffe"
	//
)

const (
	AccessKeyId     = "LTAI4FnjHfoFomah6s5vbUEm"
	AccessKeySecret = "Bp150bRDD30Tkoo9KgsmPCFG2L2g5w"
	Endpoint        = "dm.aliyuncs.com"
	EndpointDysms   = "dysmsapi.aliyuncs.com"
	EndpointVoice   = "dyvmsapi.aliyuncs.com"
)

func initProgram() error {
	//配置文件  日志  小强db  nats
	fmt.Println("Log:系统初始化中···")
	err := service.SetupSetting()
	if err != nil {
		fmt.Printf("init.setupSetting err:%v", err)
		return err
	}
	// fmt.Println("init.setupSetting success")
	// fmt.Printf("server    settings:%v\n", *global.ServerSetting)
	// fmt.Println("ali      settings:", *global.AliSetting)
	// fmt.Println("ngiot 	  settings:", *global.NgiotSetting)
	// fmt.Println("database settings:", global.DatabaseSetting)
	// fmt.Println("nats     settings:", *global.NatsSetting)
	// fmt.Println("ssdb     settings:", *global.SSDBSetting)

	//初始化日志
	service.Initlog()
	fmt.Println("init.Initlog sucess")
	//初始化数据库连接
	err = service.SetupDBEngine()
	if err != nil {
		fmt.Printf("init.setupDBEngine err:%v", err)
		return err
	}
	fmt.Println("init.setupDBEngine sucess")

	err = service.Init_nates()
	if err != nil {
		app.WriteErrorlog("init._init_nates err:  %v", err)
		return err
	}
	fmt.Println("init.Init_nates sucess")

	//初始化阿里客户端   这个客户端可多长时间会断连  断连怎么处理
	ali.AliService = ali.NewClient()
	if ali.AliService == nil {
		app.WriteErrorlog("init._init_aliClient fail")
		return err
	}
	fmt.Println("init.Init aliClient sucess")

	return nil
}

func main() {

	//初始化
	initProgram()

	cachelib.CacheInit()

	global.Recover = fmx.RecoverFn(func(szLog string) {
		// fmt.Println("panic:", szLog)
		app.WriteErrorlog("panic:%v", szLog)
	})
	//ssdb连接
	ssdbconn := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "chenxi",
		Subsystem: "elevatorServer",
		Name:      "SSDBStatus",
		Help:      "SSDB connection Status",
	})
	prometheus.MustRegister(ssdbconn)
	//cocakroach连接
	cockroachdbconn := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "socservice",
		Subsystem: "socservice",
		Name:      "cockroachdbconn",
		Help:      "db cockroachdbconn",
	})
	prometheus.MustRegister(cockroachdbconn)

	router := routers.NewRouter()
	router.GET("/prometheus/metrics", func(c *fmx.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// var t service.Topic
	// strJson := `{"data":{"events":[{"code":7002,"level":8,"paras":[],"startTime":1673516376704,"startTime_tz":480,"type":1}]},"id":"2970a7bd-bf46-4c6b-8f52-21e2a2960d24","message":"","module":"arm","sn":"E0AD12345D0000000056","ts":1673516376706}`
	// t.HandlerEventRep([]byte(strJson))
	//URORA test 极光推送
	//appMsgPusher = jpush.NewJPush()
	///appMsgPusher.JpushDemo()

	//检查连接
	go func() {
		//svc := service.New(context.TODO())
		cockroachdbconn.Set(1)
		for {
			time.Sleep(time.Second * 30)
			cockconn := global.DBEngine.NewConn()
			if cockconn == nil {
				cockroachdbconn.Set(0)
				fmt.Println("checkStatus cockroach connection error")
			} else {
				defer cockconn.Close()
				_, err := cockconn.Query("select * from bc_user LIMIT 10 OFFSET 0;")
				if err != nil {
					cockroachdbconn.Set(0)
					fmt.Printf("checkStatus cockroach error :%v", err)
				} else {
					cockroachdbconn.Set(1)
				}
			}

		}
	}()

	//SingleCallByTts()
	//SingleSendMail() //发送单封邮件测试
	//SendBatchMail() //批量发送邮件测试
	//SendBatchSmsTest() //批量发送短信测试
	//SendSmsTest() //发送单条短信测试
	//前端设置配置方式   存ssdb
	//接受服务通知异常   取ssdb   调用阿里接口通知用户
	s.ListenAndServe()
}
