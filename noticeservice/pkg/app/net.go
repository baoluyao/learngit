package app

import (
	"bytes"
	"encoding/json"
	"net"
	"strings"
	"time"

	// "ffplat/pkg/app"

	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpPost(url string, param map[string]interface{}) ([]byte, int, error) {
	jsonstr, _ := json.Marshal(param)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonstr)) //这里直接调用了http里面的get函数，传入对应url即可 返回的是http包中的respoonse类型
	if err != nil {
		return nil, 500, err
	}
	defer resp.Body.Close()                //一定要关闭返回的response中的body
	body, err := ioutil.ReadAll(resp.Body) //读取body中的信息
	if err != nil {
		return nil, 500, err
	}
	//Println(string(body))
	return body, resp.StatusCode, nil

}

func HttpMethod(method, url string, param map[string]interface{}) ([]byte, int, error) {
	jsonstr, _ := json.Marshal(param)
	client := http.Client{Timeout: time.Duration(10 * time.Second)}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonstr))
	if err != nil {
		return nil, 500, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		if hasTimedOut(err) {
			return nil, 408, nil
		}
		WriteErrorlog("HttpMethod error ：%v,param:%s", err, string(jsonstr))
		return nil, 500, err
	}
	defer resp.Body.Close()                //一定要关闭返回的response中的body
	body, err := ioutil.ReadAll(resp.Body) //读取body中的信息
	if err != nil {
		return nil, 500, err
	}
	//Println(string(body))
	return body, resp.StatusCode, nil

}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	}
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}
