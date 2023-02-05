package global

import (
	"time"

	dbserver "github.com/carr123/easysql/cockroach"
	"github.com/nats-io/nats.go"
)

var (
	//DBEngine *gorm.DB

	DBEngine *dbserver.DBServer

	Nats_client *nats.Conn

	Recover              func()
	SocUser_Token        string
	SocUser_RefreshToken string
	SocUserID            string
	SocUser_Expires_at   time.Time
)
