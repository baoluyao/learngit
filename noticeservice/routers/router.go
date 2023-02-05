package routers

import (
	"noticeservice/helper/middleware"
	"noticeservice/helper/service"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/carr123/fmx"
)

func getAppDir() string {
	file, _ := exec.LookPath(os.Args[0])
	apppath, _ := filepath.Abs(file)
	dir := filepath.Dir(apppath)
	return dir
}

func NewRouter() *fmx.Engine {
	router := fmx.NewServeMux()

	router.Use(middleware.Recovery())
	// szDir := filepath.Join(getAppDir(), "/commonservice_lc2.0")
	// router.ServeDir("/", szDir)

	//家用登录接口测试
	router.POST("/notice/api/test", service.SendSmsTest)

	//router.POST("/notice/api/test", service.SendSmsTest)
	router.POST("/notice/api/getNoticeMsg", GetNoticeMsg)
	router.POST("/notice/api/readMsg", ReadMsg)
	router.POST("/notice/api/getUserRobotsMsg", GetUserRobotsMsg)

	return router
}
