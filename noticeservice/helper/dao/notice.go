package dao

import (
	"noticeservice/model"
	"noticeservice/pkg/app"
	"noticeservice/pkg/errcode"

	"github.com/google/uuid"
)

//notice_schemes
func (d *Dao) GetContentIdsByBobotInfo(orgId, productCode string) ([]uuid.UUID, error) {

	conn := d.engine.NewConn()
	defer conn.Close()

	var contentIds []uuid.UUID

	sql := `select contentid from notice_contenthead where contentid in
	(select contentid from notice_schemes where orgid=?) and array_position(productcode,?) is not null`
	args := []interface{}{
		orgId,
		productCode,
	}

	err := conn.Select(&contentIds, sql, args...)

	if err != nil {
		app.WriteErrorlog("dao GetContentIdsByBobotInfo error:%v", err)
		return nil, err
	}

	return contentIds, nil
}

// select roleids,userids,stationmsgcodes,msgcodes,phonecodes from notice_schemes where contentid =(select contentid from notice_contenthead where contentid in
// 	(select contentid from notice_schemes where orgid='2c3f38ff-4436-4a84-8cf0-2d381b350ffe') and array_position(productcode,'seaTide') is not null)
//通知谁  通知方式
//notice_schemes
func (d *Dao) GetSchemeDetailByContentId(contentId uuid.UUID) (*model.Notice_Schemes, error) {

	conn := d.engine.NewConn()
	defer conn.Close()

	var schemes []*model.Notice_Schemes

	sql := `select * from notice_schemes where contentid=?`

	err := conn.Select(&schemes, sql, contentId)

	if err != nil {
		app.WriteErrorlog("dao GetSchemeDetailByContentId error:%v", err)
		return nil, err
	}

	if len(schemes) > 0 {
		return schemes[0], nil
	}
	return nil, nil
}

//通知内容
//notice_contentDetail
//select * from notice_contentdetail where contentid='91e6e2ac-b5a5-4d61-bd56-bbc6cc5a1d19'  and noticecode ='1002' and productcode ='seaTide'
func (d *Dao) GetContentDetail(noticeCode, productCode string, contentId uuid.UUID) (*model.Notice_ContentDetail, error) {
	conn := d.engine.NewConn()
	defer conn.Close()

	var contentDetail []*model.Notice_ContentDetail

	sql := `select * from notice_contentdetail where contentid=?  and noticecode =? and productcode =?`
	args := []interface{}{
		contentId,
		noticeCode,
		productCode,
	}

	err := conn.Select(&contentDetail, sql, args...)

	if err != nil {
		app.WriteErrorlog("dao GetContentDetail error:%v", err)
		return nil, err
	}
	if len(contentDetail) > 0 {
		return contentDetail[0], nil
	}
	//app.WriteErrorlog("dao GetContentDetail nil")
	return nil, nil
}

// select c.userid,c.username,c.email,c.phone from (select userid from re_roleuser where roleid='1f6ffc2e-ae4a-431c-977f-85ceb33cb61f') b
//left join bc_user  c on b.userid =c.userid
//roleid -> user -> username -> []phonenumber
//
func (d *Dao) GetUsersByRoleIds(roleIds model.StringArray, orgId string) ([]*model.Notice_Users, error) {
	conn := d.engine.NewConn()
	defer conn.Close()
	var roleids []string

	if len(roleIds) == 0 {
		return nil, errcode.RoleNotExist
	}
	for _, roleid := range roleIds {
		roleids = append(roleids, roleid)
	}
	var noticeUsers []*model.Notice_Users

	//判断条件需要加orgid
	sql := `select c.userid,c.username,c.email,c.phone from (select userid from re_roleuser where roleid in(?) and roleid in(select cast(roleid as string) from ims_role where orgid=?)) b
            left join bc_user  c on b.userid =c.userid and c.phone is not null`

	err := conn.Select(&noticeUsers, sql, roleids, orgId)
	if err != nil {
		app.WriteErrorlog("dao GetUsersByRoleIds error:%v", err)
		return nil, err
	}

	if len(noticeUsers) > 0 {
		return noticeUsers, nil
	}
	//app.WriteErrorlog("dao GetUsersByRoleIds nil")
	return nil, nil

}

func (d *Dao) GetNoticeChannels(noticeCode string, contentId uuid.UUID) (*model.NoticeChannel, error) {
	conn := d.engine.NewConn()
	defer conn.Close()

	var noticeChannels []*model.NoticeChannel

	sql := `select  array_position(stationmsgcodes,?) as stationmsg,array_position(msgcodes,?) as msg,array_position(phonecodes,?) as phone  
	from notice_schemes where contentid = ?`

	args := []interface{}{
		noticeCode,
		noticeCode,
		noticeCode,
		contentId,
	}
	err := conn.Select(&noticeChannels, sql, args...)
	if err != nil {
		app.WriteErrorlog("dao GetNoticeChannels error:%v", err)
		return nil, err
	}
	if len(noticeChannels) > 0 {
		return noticeChannels[0], nil
	}

	return nil, nil
}
