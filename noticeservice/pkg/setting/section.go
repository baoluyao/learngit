package setting

import (
	"errors"
	"time"
)

type ServerSettings struct {
	ServerName     string
	LogKeep        int
	HeartBeat      int
	HttpPort       string
	Https          bool
	Timeout        time.Duration
	LogRecord      bool
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	LogSavePath    string
	LogFileSize    int64
	LogFileCount   int64
	LogLevel       int
	LogMode        int
	FileCleanTime  int
	FileExpireTime int
}

type AliSettings struct {
	AccessKeyId          string
	AccessKeySecret      string
	AccessKeyIdVoice     string
	AccessKeySecretVoice string
	EndpointDm           string
	EndpointDysms        string
	EndpointVoice        string
}

type NgiotSettings struct {
	Addr            string
	BaseUrl         string
	ClientId        string
	ClientSecret    string
	User            string
	Password        string
	Resource        string
	Country         string
	Opt             string
	ChallengeType   string
	ChallengeSecret string
}

type DatabaseSettings struct {
	Addr         string
	MaxIdleConns int
}

type NatsSettings struct {
	Connect             string
	Name                string
	Timeout             time.Duration
	PingInterval        time.Duration
	MaxPingsOutstanding int
	MaxReconnect        int
	ReconnectWait       time.Duration
	ReconnectBufSize    int
	UserName            string
	Password            string
}

type SSDBSettings struct {
	Addr         string
	Addr1        string
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if v == nil {
		return errors.New("ReadSection " + k + " nil")
	}
	return nil
}
