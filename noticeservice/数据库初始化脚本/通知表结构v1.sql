-- 通知项配置



INSERT INTO public.notice_items (itemid, errorcode, productid, "type", "level", remark, "module", title, content, solution, jumplink, createtime, createuser, updatetime, updateuser) VALUES
('aba8bb5c-2fcf-46d5-931f-3f3cc8de4aff', 1001, '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', 1, 1, '', 1, '地磁触发', '机器处于地磁触发状态,请手动解除。', '', '', 0, 'admin', 0, ''),
('8470f71e-cc50-4323-99c2-cd96fe02531e', 1002, '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', 1, 1, '', 1, '机器发生碰撞', '机器发生碰撞，停止移动，请帮助远离碰撞障碍物，并检查碰撞物体是否有损坏。', '', '', 0, 'admin', 0, ''),
('4639f379-41d3-4a67-93dd-506249f08514', 1003, '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', 1, 1, '', 2, '机器打滑', '机器当前行驶区域地面光滑，请手动帮助机器远离此区域。', '', '', 0, 'admin', 0, ''),
('6fb8d6bc-0e0f-4068-bfbd-66c202c5d851', 1004, '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', 1, 1, '', 2, '机器倾斜', '机器当前处于倾斜状态，请帮助机器恢复正常.', '', '', 0, 'admin', 0, ''),
('68cf0f65-069c-4aeb-81c4-1aae2247ab2a', 1005, '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', 1, 1, '', 2, '前滚刷未安装', '请先安装推尘布滚刷，再启动工作', '', '跳转：对应安装说明页面', 0, 'admin', 0, ''),
('139daa0a-0980-4013-8848-221146aeab1d', 1006, '122f24d9-5951-4c39-93d0-4e8f38b41e34', 1, 1, '', 3, '梯控通信异常', '请点击查看解决方案', '解决方案：
1. 请重启机器和充电桩Lora
2. 若仍无法解决，请拨打XXXX（售后电话，后台可配置）联系售后维修检查', '', 0, 'admin', 0, ''),
('79a933d8-d0ea-4f6f-a0a7-07d8914e2daa', 1007, '122f24d9-5951-4c39-93d0-4e8f38b41e34', 2, 1, '', 1, '急停触发', '请手动解除机器急停状态', '', '', 0, 'admin', 0, ''),
('00269ec9-a41c-4eae-a104-7be5d6b601dc', 1008, '122f24d9-5951-4c39-93d0-4e8f38b41e34', 2, 1, '', 5, '无法找到充电桩或补给站', '请在地图上创建充电桩或补给站位置', '', '跳转：地图编辑-点位管理页面', 0, 'admin', 0, ''),
('b71f0d2c-f45e-4d4c-83f3-185c17c37fa0', 1009, '122f24d9-5951-4c39-93d0-4e8f38b41e34', 2, 1, '', 2, '头盖未合盖', '请先关闭头盖，再启动工作', '', '', 0, 'admin', 0, ''),
('852dd667-dcde-4f6f-a03d-566cee87b722', 1010, '122f24d9-5951-4c39-93d0-4e8f38b41e34', 2, 1, '', 8, '垃圾盒过滤网重度损耗', '为保证清洁效果，请立即更换垃圾盒过滤网，点击查看详情。', '', '跳转：对应耗材管理页面（按照更换HEPA操作步骤来呈现）', 0, 'admin', 0, '');

-- 机构  orgid  d057677b-d497-45d1-8df6-793fa8cdeae1
-- 产品类型 productid  (海浪)21fe2142-151d-4cdd-8fd4-f8b17e24f2fc  (海潮)5f854021-c19c-40e3-a693-2fb55cc36ab9   (T9)28758389-c9a8-4b85-807f-e5d77cf0a9ec

-- 事件类型
-- 故障异常： 1     状态提醒 ：2

-- 模块
-- 整机状态： 1     耗材：8
-- 硬件组件： 2     工作：9
-- 电器组件： 3     地图：10
-- 传感器：   4     升级：11
-- 软件系统： 5     移动流量：12
-- 补给站：   6     地图编辑：13
-- 充电桩：   7

-- 通知方案配置
--统一机构下的产品类型只有一种方案
INSERT INTO public.notice_schemes (schemeid, orgid, productid, roleids, userids, stationmsgcodes, msgcodes, phonecodes, createtime, createuser, updatetime, updateuser) VALUES
('c6aae845-3786-4ca3-8884-b4e6d3e6990b', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', '{09587a6d-1f92-42da-9051-0cfcb4c5bd9f}', '{}', '{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),
('3ca6d851-e0d5-4db2-9d06-8c24c92cfac9', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', '{0d7f3873-f077-48c2-91a1-d8ff544b98d3}', '{}','{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),
('07a09a8c-11e3-40b0-b400-8ec52c95cc0f', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', '{893b29ca-8ba2-46dc-8160-6a5328b9a80e}', '{}', '{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),
('b9de567e-b7a4-4f6c-bdcc-ee4a67a2871d', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '5f854021-c19c-40e3-a693-2fb55cc36ab9', '{d057677b-d497-45d1-8df6-793fa8cdeae1}', '{}','{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),
('02112b97-ce17-40e2-80f1-0dbe66f1fd5f', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '5f854021-c19c-40e3-a693-2fb55cc36ab9', '{09587a6d-1f92-42da-9051-0cfcb4c5bd9f,0d7f3873-f077-48c2-91a1-d8ff544b98d3}', '{}', '{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),
('4237d9cf-805c-4bde-96d9-9a29e460fd4e', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '5f854021-c19c-40e3-a693-2fb55cc36ab9', '{893b29ca-8ba2-46dc-8160-6a5328b9a80e,d057677b-d497-45d1-8df6-793fa8cdeae1}', '{}', '{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),

d057677b-d497-45d1-8df6-793fa8cdeae1
-- robotsn     orgid productid errorcode
-- 查找配置方案 orgid productid
-- select * notice_shemes where orgid='{}' and productid='{}' and stationmsg
-- 哪些codes存在 改errorcode 



drop table nt_noticedefault;
drop table nt_noticedefaultdetail;
drop table nt_noticedetail;
drop table nt_noticehead;

 DROP TABLE public.notice_items;

--  状态配置项
CREATE TABLE IF NOT EXISTS public.notice_items (
	itemid uuid not null,
	errorcode text null,
	productid uuid null,
	type int null,
	level int null,
	remark text null,
	module int null,
	title text null,
	content text null,
	solution text null,
	jumplink text null,
	createtime int null,
	createuser text null,
	updatetime int null,
	updateuser text null,
	CONSTRAINT noticeitems_pkey PRIMARY KEY (itemid ASC)
);

--通知配置方案
CREATE TABLE IF NOT EXISTS public.notice_schemes (
	schemeid uuid not null,
	orgid uuid not null,
	productid uuid null,
	roleids _text null,
	userids _text null,
	stationmsgcodes _text null,
	msgcodes _text null,
	phonecodes _text null,
	createtime int null,
	createuser text null,
	updatetime int null,
	updateuser text null,
	CONSTRAINT noticeschemes_pkey PRIMARY KEY (schemeid ASC)
);

--通知记录 主表
CREATE TABLE IF NOT EXISTS public.notice_recordhead (
	recordid uuid not null,
	orgid uuid not null,
	productid uuid  not null,
	noticeType int null,
	robotsn text not null,
	errorcode text not null,
	userids  _text null,
	noticetime int null,
	CONSTRAINT nt_recordhead_pkey PRIMARY KEY (recordid ASC)
);

--通知记录明细
CREATE TABLE IF NOT EXISTS public.notice_recorddetail (
	recordid uuid not null,
	userid uuid null,
	phoneNumber text not null,
	channel int not null,
	result int null
);




