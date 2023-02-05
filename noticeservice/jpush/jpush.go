package jpush

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"noticeservice/model"

	"github.com/carr123/fmx"
)

type PushMsgAndroid1 struct {
	//Platform []string `json:"platform"`
	Platform string `json:"platform"`
	Audience struct {
		Alias []string `json:"alias"`
	} `json:"audience"`
	Notification struct {
		Alert   string `json:"alert"`
		Android struct {
			Title string `json:"title"`
			//Priority   int                    `json:"priority"`   //通知栏展示优先级
			//Builder_id int                    `json:"builder_id"` //通知栏样式 ID
			Extras map[string]interface{} `json:"extras"`
		} `json:"android"`
	} `json:"notification"`

	Message struct {
		Content     string `json:"msg_content"`
		ContentType string `json:"content_type"`
		Title       string `json:"title"`
	} `json:"message"`
}

func NewPushMsgAndroid1(alias []string, Title string, Alert string, message string, Priority int, Builder_id int, Extras map[string]interface{}) (*PushMsgAndroid1, error) {
	if len(alias) > 1000 {
		return nil, errors.New("alias too much")
	}

	msg := &PushMsgAndroid1{}
	//msg.Platform = []string{"android"}

	//msg.Platform = []string{"all"}
	msg.Platform = "all"
	msg.Audience.Alias = make([]string, 0, len(alias))
	msg.Audience.Alias = append(msg.Audience.Alias, alias...)
	msg.Notification.Alert = Alert         //通知内容   eg:  机器正在进行开机自检
	msg.Notification.Android.Title = Title //通知标题   eg:  1-自检

	//安卓端不需要这个了
	//msg.Notification.Android.Priority = Priority
	//msg.Notification.Android.Builder_id = Builder_id

	if Extras != nil {
		msg.Notification.Android.Extras = Extras
	} else {
		msg.Notification.Android.Extras = make(map[string]interface{})
	}
	msg.Message.Content = message
	msg.Message.ContentType = "text"
	msg.Message.Title = Title

	return msg, nil
}

type JPUSH struct {
	appKey       string
	masterSecret string
	pushURL      string
}

func NewJPush() *JPUSH {
	return &JPUSH{}
}

func (t *JPUSH) SetURL(pushURL string) {
	t.pushURL = pushURL //"https://api.jpush.cn/v3/push", ///validate
}

func (t *JPUSH) SetAuth(appKey string, masterSecret string) {
	t.appKey = appKey
	t.masterSecret = masterSecret
}

func (t *JPUSH) DoPush(msg interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(msg)
	if err != nil {
		return nil, fmx.Error(err)
	}

	req, err := http.NewRequest("POST", t.pushURL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmx.Error(err)
	}

	req.SetBasicAuth(t.appKey, t.masterSecret)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmx.Error(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmx.Error(err)
	}

	if resp.StatusCode != 200 {
		return body, fmx.Errorf(500, "jpush respcode:%d, resp:%s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (t *JPUSH) JpushDemo(event *model.Event) {
	// svc:=service.New{}

	var pusher *JPUSH = NewJPush()
	appKey := "3e899a7a7267e446d9531464"
	masterSecret := "a8d84afae6665acd334a8c64"
	//past appKey、masterSecret
	// appKey := "3c419c716a46d165f98b8f3d"
	// masterSecret := "c2afeffc2dbe50b28c4dcecd"
	pusher.SetURL("https://api.jpush.cn/v3/push")
	pusher.SetAuth(appKey, masterSecret)

	alias := []string{"luyanqi"} //useraccount
	//alias := []string{"luyanqi1"} //useraccount
	//alias []string, Title string, Alert string, message string, Priority int, Builder_id int, Extras map[string]interface{}
	extras := map[string]interface{}{
		// "code": event.Code,
		// "type": event.Code,
		// "code": event.Code,
		// "code": event.Code,
		// "code": event.Code,
		// "code": event.Code,
		"event": event,
	}
	//alert 通知内容   title  通知标题
	msg1, _ := NewPushMsgAndroid1(alias, "开机自检", "机器人正在进行开机自检！", "机器人正在进行开机自检！", 1, 2, extras)
	//msg1, _ := NewPushMsgAndroid1(alias, "title", "alert", "hello world", 1, 2, nil)

	res, err := pusher.DoPush(msg1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(res))
	//{"sendno":"0","msg_id":"18100528814460569"}

}
