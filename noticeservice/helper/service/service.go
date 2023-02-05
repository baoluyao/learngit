package service

import (
	"context"
	"noticeservice/global"
	"noticeservice/helper/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	// svc := Service{ctx: ctx}
	// svc.dao = dao.New(global.DBEngine)
	svc := Service{ctx: ctx, dao: dao.New(global.DBEngine)}
	return svc
}
