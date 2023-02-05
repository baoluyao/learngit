package dao

import (
	dbserver "github.com/carr123/easysql/cockroach"
)

type Dao struct {
	engine *dbserver.DBServer
}

func New(_engine *dbserver.DBServer) *Dao {
	return &Dao{engine: _engine}
}
