package setting

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	// //获取程序实际路径
	// _, path := getAppPath()
	// fullPath := filepath.Join(path, "configs")
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	// vp.AddConfigPath(fullPath)
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func getAppPath() (string, string) {
	file, _ := exec.LookPath(os.Args[0])
	apppath, _ := filepath.Abs(file)
	dir := filepath.Dir(apppath)
	return apppath, dir
}
