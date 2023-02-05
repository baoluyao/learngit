 DROP TABLE IF EXISTS public.notice_items;

 --  状态配置项
CREATE TABLE IF NOT EXISTS public.notice_items (
	itemid uuid not null,
    productcode text not null, --产品类型
	noticecode text null,      --状态码
    notice   text null,   -- 状态描述？
	noticetypeid uuid not null, --一级功能模块id
    noticeType  text null,
	level int null,   --状态等级（0 状态 1 异常 2 告警）
    scencedesc text null,  --场景描述
	createtime int null, --创建时间 
	createuser text null, --创建人
	updatetime int null,  --更新时间
	updateuser text null,  --更新人
	CONSTRAINT noticeitems_pkey PRIMARY KEY (itemid ASC)
);

 DROP TABLE IF EXISTS public.notice_contenthead;
-- 通知内容方案表 notice_contenthead
CREATE TABLE IF NOT EXISTS public.notice_contenthead (
    contentid uuid not null,
    contentName text not null,
    productcode  _text null,
    createtime int null, --创建时间 
	createuser text null, --创建人
	updatetime int null,  --更新时间
	updateuser text null,  --更新人
    CONSTRAINT noticecontenthead_pkey PRIMARY KEY (contentid ASC)
);

 DROP TABLE IF EXISTS public.notice_contentdetail;
-- 通知方案明细表 notice_contentdetail
CREATE TABLE IF NOT EXISTS public.notice_contentdetail (
    contentid uuid not null,    --内容方案id
    productcode text not null,  --产品类型code 
    noticecode text not  null,  --通知 状态码
    title text not null,        --通知 标题
    content text not null,      --通知内容
    solution text null,         --解决方案
    jumplink text null          --跳转链接
);


DROP TABLE IF EXISTS public.notice_schemes;
-- 通知配置方案   notice_schemes
CREATE TABLE IF NOT EXISTS public.notice_schemes (
	schemeid uuid not null,
	orgid uuid not null,
	productcode uuid not null,
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

-- 通知项

-- v2
INSERT INTO public.notice_items (itemid, productcode, noticecode, "notice", noticetypeid, noticetype, "level", content, createtime, createuser, updatetime, updateuser) VALUES
('aba8bb5c-2fcf-46d5-931f-3f3cc8de4aff','seaTide' ,1001, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'地磁触发',0, 'admin', 0, ''),
('8470f71e-cc50-4323-99c2-cd96fe02531e','seaTide' ,1002, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'机器发生碰撞',0, 'admin', 0, ''),
('4639f379-41d3-4a67-93dd-506249f08514','seaWave' ,1003, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'前滚刷未安装', ,0, 'admin', 0, ''),
('6fb8d6bc-0e0f-4068-bfbd-66c202c5d851','seaWave' ,1004, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'梯控通信异常', ,0, 'admin', 0, ''),

('68cf0f65-069c-4aeb-81c4-1aae2247ab2a','seaTide' ,1005, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'机器打滑', 0, 'admin', 0, ''),
('139daa0a-0980-4013-8848-221146aeab1d','seaTide' ,1006, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'机器倾斜', 0, 'admin', 0, ''),
('79a933d8-d0ea-4f6f-a0a7-07d8914e2daa','seaWave' ,1007, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'急停触发', 0, 'admin', 0, ''),
('00269ec9-a41c-4eae-a104-7be5d6b601dc','seaWave' ,1008, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'头盖未合盖', 0, 'admin', 0, ''),

('b71f0d2c-f45e-4d4c-83f3-185c17c37fa0','seaTide' ,1009, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'地磁触发', 0, 'admin', 0, ''),
('852dd667-dcde-4f6f-a03d-566cee87b722','seaTide' ,1010, '状态描述','39a0eaaf-c846-4bc0-8acb-2f54ef7b66dd','一级功能模块描述',0,'地磁触发', 0, 'admin', 0, ''),

-- v1
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

-- 内容方案
INSERT INTO public.notice_contenthead (contentid, contentname, productcode, createtime, createuser, updatetime, updateuser) VALUES
('91e6e2ac-b5a5-4d61-bd56-bbc6cc5a1d19', '方案1', '{seaTide,seaWave}', 1666002814056, 'admin', 0, ''),
('db7d0e3d-b531-4802-a412-ac38c2347148', '方案2', '{seaTide,seaWave}', 1666002814056, 'admin', 0, ''),
('f3270535-e925-4649-ab0d-bfb0e53994cd', '方案3',  '{seaTide}', 1666002814056, 'admin', 0, ''),
('36901d1b-b530-429e-90bc-ebc69a1c805d', '方案4',  '{seaTide}', 1666002814056, 'admin', 0, ''),
('f4d87146-8b58-495b-b6c9-afc0d6e9f9ad', '方案5',  '{seaWave}', 1666002814056, 'admin', 0, ''),
('9d4a5776-fbaa-4317-b00f-67e88fec2045', '方案6',  '{seaWave}', 1666002814056, 'admin', 0, ''),



-- 内容明细
INSERT INTO public.notice_contentdetail (contentid, productcode, noticecode, title, content, solution, jumplink) VALUES
('91e6e2ac-b5a5-4d61-bd56-bbc6cc5a1d19', 'seaTide', '1001', '地磁触发', '机器处于地磁触发状态,请手动解除。', '', ''),
('91e6e2ac-b5a5-4d61-bd56-bbc6cc5a1d19', 'seaTide', '1002', '机器发生碰撞', '机器发生碰撞，停止移动，请帮助远离碰撞障碍物，并检查碰撞物体是否有损坏。', '', ''),
('91e6e2ac-b5a5-4d61-bd56-bbc6cc5a1d19', 'seaWave', '1003', '前滚刷未安装', '请先安装推尘布滚刷，再启动工作', '', ''),
('91e6e2ac-b5a5-4d61-bd56-bbc6cc5a1d19', 'seaWave', '1004', '梯控通信异常', '请点击查看解决方案', '', ''),


-- 通知配置

INSERT INTO public.notice_schemes (schemeid, orgid, roleids, userids, stationmsgcodes, msgcodes, phonecodes, createtime, createuser, updatetime, updateuser) VALUES
(?, ?, ?, ?, ?, ?, ?, ?, 0, '', 0, '');


INSERT INTO public.notice_schemes (schemeid, orgid, roleids, userids, stationmsgcodes, msgcodes, phonecodes, createtime, createuser, updatetime, updateuser) VALUES
('c6aae845-3786-4ca3-8884-b4e6d3e6990b', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '{09587a6d-1f92-42da-9051-0cfcb4c5bd9f}', '{}', '{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),

('3ca6d851-e0d5-4db2-9d06-8c24c92cfac9', '2c3f38ff-4436-4a84-8cf0-2d381b350ffe', '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', '{0d7f3873-f077-48c2-91a1-d8ff544b98d3}', '{}','{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),
('07a09a8c-11e3-40b0-b400-8ec52c95cc0f', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '21fe2142-151d-4cdd-8fd4-f8b17e24f2fc', '{893b29ca-8ba2-46dc-8160-6a5328b9a80e}', '{}', '{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),

('b9de567e-b7a4-4f6c-bdcc-ee4a67a2871d', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '5f854021-c19c-40e3-a693-2fb55cc36ab9', '{d057677b-d497-45d1-8df6-793fa8cdeae1}', '{}','{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),
('02112b97-ce17-40e2-80f1-0dbe66f1fd5f', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '5f854021-c19c-40e3-a693-2fb55cc36ab9', '{09587a6d-1f92-42da-9051-0cfcb4c5bd9f,0d7f3873-f077-48c2-91a1-d8ff544b98d3}', '{}', '{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),
('4237d9cf-805c-4bde-96d9-9a29e460fd4e', 'd057677b-d497-45d1-8df6-793fa8cdeae1', '5f854021-c19c-40e3-a693-2fb55cc36ab9', '{893b29ca-8ba2-46dc-8160-6a5328b9a80e,d057677b-d497-45d1-8df6-793fa8cdeae1}', '{}', '{1001}', '{1001,1002}', '{1001,1002,1003}', 0, 'admin', 0, ''),




-- 通知表结构v2.sql  变更内容
-- 1.配置项表   场景内容content  变为 scencedesc场景描述
-- 2.
-- 3.通知配置表  拿掉了 productcode 产品类型