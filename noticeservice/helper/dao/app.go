package dao

import (
	"noticeservice/model"
	"noticeservice/pkg/app"
	"noticeservice/pkg/errcode"
	"strconv"
)

func (d *Dao) GetNoticeMsg(param *model.GetNoticeMsg) ([]*model.AppNoticeMsg, *model.Pager, *errcode.Error) {
	var pageCond string

	conn := d.engine.NewConn()
	defer conn.Close()

	var noticeMsg []*model.AppNoticeMsg

	sql := `select a.mid as msgid,b.noticetype,b.title,b.content,b.jumplink,a.readstatus,a.noticetime,a.noticetime_tz  from notice_recorddetail a left join notice_item b
	on a.noticecode=b.noticecode  where a.recordid in  (select recordid from notice_recordhead where robotsn=?)`
	args := []interface{}{}

	err := conn.Select(&noticeMsg, sql, args...)

	if err != nil {
		app.WriteErrorlog("dao GetContentDetail error:%v", err)
		return nil, param.Page, errcode.DBOperationError
	}
	if param.Page != nil {
		if rows, err := d.GetNoticeMsgCount(param.RobotId); err != nil {
			return nil, param.Page, err
		} else {
			param.Page.TotalRows, _ = strconv.Atoi(strconv.FormatInt(rows, 10))
		}
		pageCond += " LIMIT ?	OFFSET ?"
		args = append(args, param.Page.PageSize)
		args = append(args, param.Page.PageSize*(param.Page.CurPage-1))
	}
	return nil, param.Page, nil
}

func (d *Dao) GetNoticeMsgCount(robotId string) (int64, *errcode.Error) {
	conn := d.engine.NewConn()
	defer conn.Close()

	sql := `select count(mid) from notice_recorddetail  where 
	recordid in (select recordid from notice_recordhead where robotsn= ? )`

	num, err := conn.QueryCount(sql, robotId)
	if err != nil {
		app.WriteErrorlog("dao.GetNoticeMsgCount error: %v,sql:%s,paras:%v", err, sql, robotId)
		return 0, errcode.DBOperationError
	}

	return num, nil
}

func (d *Dao) GetRobotSNByDid(did []string) ([]string, *errcode.Error) {
	conn := d.engine.NewConn()
	defer conn.Close()

	var robotSN []string
	sql := `select sn from ngiot_info where ngiot_did in (?) `

	if err := conn.Select(&robotSN, sql, did); err != nil {
		return nil, errcode.DBOperationError
	}

	return robotSN, nil

}

func (d *Dao) GetRobotsMsg(robotId []string) (*[]model.RobotMsgList, *errcode.Error) {
	conn := d.engine.NewConn()
	defer conn.Close()

	var robotMsgList *[]model.RobotMsgList
	sql := ` select robotid,orgid,totalcount,y.unreadcount from (select b.recordid as recordid,robotsn as robotid,orgid,count(b.recordid) as totalcount from notice_recordhead a left join 
	notice_recorddetail b  on a.recordid = b.recordid 
	and a.robotsn in(?) group by robotsn,orgid,b.recordid)x left join ( select recordid,count(readstatus) as unreadcount from notice_recorddetail where  readstatus=false group by recordid  ) y
	on x.recordid = y.recordid
	 `

	if err := conn.Select(&robotMsgList, sql, robotId); err != nil {
		return nil, errcode.DBOperationError
	}

	return robotMsgList, nil

}

func (d *Dao) ReadAllMsg(robotId string) *errcode.Error {
	conn := d.engine.NewConn()
	defer conn.Close()

	sql := `update  notice_recorddetail  set readstatus = true where recordid in
	(select recordid from notice_recordhead where robotsn=?) `

	if err := conn.Exec(sql, robotId); err != nil {
		return errcode.DBOperationError
	}

	return nil
}

//#region

func (d *Dao) ReadSingleMsg(msgId string) *errcode.Error {
	conn := d.engine.NewConn()
	defer conn.Close()

	sql := `update  notice_recorddetail  set readstatus = true where mid= ? `

	if err := conn.Exec(sql, msgId); err != nil {
		return errcode.DBOperationError
	}

	return nil
}

//#endregion
