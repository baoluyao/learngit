package errcode

//行备注statuscode 返回值
//若无备注，默认200
var (
	//梯控http请求 0成功  1失败

	//预约机器人
	RoleNotExist      = NewError(1001, "role not exist")                 //400
	GetNoticeInfoFail = NewError(1002, "users or contentdetail is null") //400

	Success          = NewError(100000000, "success")                          //200
	ServerError      = NewError(100000001, "server error")                     //500
	DBOperationError = NewError(100000002, "sql operation exception")          //500
	InvalidParams    = NewError(100000003, "invalid params")                   //400
	NgIOTServerError = NewError(100000004, "ngiot server account login error") //500

	UnauthorizedAuthNotExist = NewError(110000003, "invalid user")              //401
	UnauthorizedPlatNotExist = NewError(110000004, "unauthorized for platform") //401
	UnauthorizedUserNotExist = NewError(110000005, "user not exist")            //401
	UnauthorizedUserDisabled = NewError(110000006, "user disabled")             //401
	UnauthorizedRoleNotExist = NewError(110000007, "user have no role")         //401
	UnauthorizedTokenError   = NewError(110000008, "token error")               //401
	UnauthorizedTimeout      = NewError(110000009, "login timeout")             //401
	HaveNewToken             = NewError(110000010, "login by another address")  //401
	UnauthorizedGenerate     = NewError(110000011, "get token error")           //401
	TooManyRequests          = NewError(110000012, "too many requests")         //429

	NetError         = NewError(110000013, "network error")           //400
	RegisterNetError = NewError(110000015, "register connect error")  //400
	FilerNetError    = NewError(110000016, "file connect error")      //400
	RPCNetError      = NewError(110000017, "rpc connect error")       //400
	NbIOTNetError    = NewError(110000018, "iotserver connect error") //400
	NgIOTNetError    = NewError(110000019, "ngiot connect error")     //400

	NotFound = NewError(110000014, "data not found") //200
	//机构相关提示
	OrgNotExist         = NewError(111000001, "organization not exist")
	OrgAlreadyExist     = NewError(111000002, "org already exist")
	ParentOrgNotExist   = NewError(111000003, "parent organization not exist")
	OrgUserQuotaError   = NewError(111000004, "organization user quota error")
	OrgRobotQuotaError  = NewError(111000005, "organization robot quota error")
	OrgTypeAlreadyUsed  = NewError(111000006, "organization type already used")
	OrgTypeAlreadyExist = NewError(111000007, "organization type already exsit")
	OrgTypeTagDuplicate = NewError(111000008, "organization tag already exist")
	OrgHasRobotUsed     = NewError(111000009, "organization have robot used")
	OrgProductHasUsed   = NewError(111000010, "organization's product have robot used")
	OrgTypeNotExist     = NewError(111000011, "organization type not exsit")

	//用户相关
	UserNotExist            = NewError(112000001, "user not exist")
	UseraccountAlreadyExist = NewError(112000002, "useraccount already exist")
	UserCountOverQuata      = NewError(112000003, "user count over quota")
	UserPasswordError       = NewError(112000004, "user password error")
	UserBaseInsertNoUser    = NewError(112000005, "userbase insert user error")
	UserBaseInitAdminEroor  = NewError(112000006, "userbase init admin error")
	UserBaseError           = NewError(112000007, "userbase error")
	UserBaseClientError     = NewError(112000008, "userbase client error")
	UserBaseAuthorizeError  = NewError(112000009, "userbase authorize error")
	//role
	RoleNameAlreadyExist = NewError(113000001, "rolename already exist")
	RoleHaveNoUser       = NewError(113000002, "role have no user")
	ImsRoleNotExist      = NewError(113000003, "imsrole not exist")
	//plat
	PlatNotExist     = NewError(114000001, "platform not exist")
	PlatAlreadyExist = NewError(114000002, "platform already exist")

	//apk
	ApkNotExist      = NewError(115000001, "apk not exist")
	ApkAlreadyExist  = NewError(115000002, "apk already exist")
	ApkMissingPakage = NewError(115000003, "apk missing package")

	//version
	VersionNotExist     = NewError(116000001, "version not exist")
	VersionAlreadyUsed  = NewError(116000002, "version already used")
	VersionAlreadyExist = NewError(116000003, "version already exist")

	///robot
	RobotNotExist           = NewError(117000001, "robot not exist")
	InvalidActiveCode       = NewError(117000002, "invalid activecode")
	InvalidUserForRobot     = NewError(117000003, "invalid user for robot")
	RobotLocked             = NewError(117000004, "robot has been locked")
	RobotInvalidNgiot       = NewError(117000005, "robot invalid ngiot")
	GetRobotWorkStatusError = NewError(117000006, "get robot workstatus error")
	RobotCountOverQuota     = NewError(117000007, "robot count over quota")
	RobotLockedError        = NewError(117000008, "robot locked error")
	RobotUnLockedError      = NewError(117000009, "robot unlocked error")
	RobotSNAlreadyExist     = NewError(117000010, "robotsn already exist")
	RobotNameAlreadyExist   = NewError(117000011, "robot name already exist")
	RobotAlreadyExpired     = NewError(117000012, "robot already exist expired")

	RobotImportError      = NewError(117000013, "robot import error")
	RobotDataInvalid      = NewError(117000014, "robot data is invalid")
	RobotSNInvalid        = NewError(117000015, "robotsn is invalid")
	RobotNameInvalid      = NewError(117000016, "robot name is invalid")
	RobotDataNull         = NewError(117000017, "robot is null")
	RobotDataTooLong      = NewError(117000018, "robot data too long")
	RobotSNOnlyNumber     = NewError(117000019, "robot sn only number")
	DeleteRobotHaveLocked = NewError(117000020, "delete robot have locked")

	//file
	FileNotExist          = NewError(118000001, "file not exist")
	FileAlreadyExistInorg = NewError(118000002, "file already in this organazation")
	FileUplodFail         = NewError(118000003, "file upload fail")
	InvalidFileMd5        = NewError(118000005, "invalid file md5")
	InvalidExcelFile      = NewError(118000006, "invalid excel file")

	//map

	MapNotExist         = NewError(119000001, "map not exist")
	MapAlreadyExist     = NewError(119000002, "map file already exist")
	MapPointAlreayExist = NewError(119000003, "map point file already exist")
	MapWallAlreadyExist = NewError(119000004, "map virtual wall file already exist")
	MapRecordNotExist   = NewError(170000005, "map record not exist")
	//common
	UpdateRecordNotExist = NewError(120000001, "update record not exist")
	GetMessageError      = NewError(120000002, "get message error")
	SetMessageError      = NewError(120000003, "set message error")
	DelMessageError      = NewError(120000004, "delete message error")
	//params
	ParamNotExist        = NewError(121000001, "param not exist")
	GetRobotParamsError  = NewError(121000002, "get robot param error")
	ParaCodeAlreadyExist = NewError(121000003, "para code already exist")
	ParaNameAlreadyExist = NewError(121000004, "para name already exist")

	//module
	ModuleNotExist         = NewError(122000001, "module not exist")
	ModuleCodeAlreadyExist = NewError(122000002, "module code already exist")
	ModuleNameAlreadyExist = NewError(122000003, "module name already exist")
	ModuleAlreadyUsed      = NewError(122000004, "module already used")

	//elevator
	GetElevatorError = NewError(123000001, "get elevator error")

	//map
	GetMapPackageError    = NewError(124000001, "get map package error")
	UpdateMapPackageError = NewError(124000002, "update map package error")
	DelMapPackageError    = NewError(124000003, "delete map package error")

	ProductNotExist         = NewError(125000001, "product not exist")
	ProductCodeAlreadyExist = NewError(125000002, "product code already exist")
	ProductNameAlreadyExist = NewError(125000003, "product name already exist")
	ProductTypeNotExist     = NewError(125000004, "product type not exist")
	ProductAlreadyUsed      = NewError(125000005, "product alreay used")
	ProductTypeAlreadyUsed  = NewError(125000006, "product type alreay used")
	ProductypeAlreadyExist  = NewError(125000007, "product type already exist")
	ParentProTypeNotExist   = NewError(125000008, "parent product type not exist")
	//ngiot
	NgiotError               = NewError(126000001, "ngiot get error")
	NgiotGetShareCodeError   = NewError(126000002, "ngiot get sharecode error")
	NgiotBindRobotError      = NewError(126000003, "ngiot bind robot error")
	NgiotRobotResError       = NewError(126000004, "ngiot robot response error")
	NgiotGetBindDevicesError = NewError(126000005, "ngiot get user bind devices error")
	NgiotGetBindUsersError   = NewError(126000006, "ngiot get robot bind users error")
	NgiotUnBindDevicesError  = NewError(126000007, "ngiot unbind devices error")
	NgiotRobotNotOffline     = NewError(4200, "robot offline in ngiot") //126000008
	NgiotRobotNoBindServer   = NewError(126000009, "robot have not bind with server account")

	//nbiot
	NbiotCreateDeviceError = NewError(127000001, "nbiot create device error")
	NbiotDeleteDeviceError = NewError(127000002, "nbiot delete device error")
	NbiotQueryDeviceError  = NewError(127000003, "nbiot query device error")
	NbiotDisabled          = NewError(127000004, "nbiot disabled")
	NbiotImeiRepNull       = NewError(127000005, "nbiot report imei null")
	NbiotImeiRepException  = NewError(127000006, "nbiot report imei exception")
	NbiotSendCmdError      = NewError(127000007, "nbiot send cmd error")

	//extra_service
	ExtraServiceAlreadyExsit = NewError(128000001, "service already exist")
	ExtraServiceAlreadyOpen  = NewError(128000002, "service already open")
	ExtraServiceAlreadyClose = NewError(128000003, "service already close")

	//business clean
	GetCleanRecordError     = NewError(130000001, "get clean record error")
	GetCleanStatisticsError = NewError(130000002, "get clean statistics error")

	//business cruise
	LineNameAlreadyExist   = NewError(140000001, "lineName already exist")
	AdvertNameAlreadyExist = NewError(140000002, "advert name already exist")
	FormNameAlreadyExist   = NewError(140000003, "form name already exist")

	BuildingNameAlreadyExist  = NewError(150000001, "buiding name already exist")
	BuildingNotExist          = NewError(150000002, "buiding not exist")
	BuildingFloorAlreadyExist = NewError(150000003, "buiding floor already exist")
	RobotGroupNotExist        = NewError(160000001, "robot group not exist")
	RobotGroupAlreadytExist   = NewError(160000002, "robot group name already exist")

	LogRecordNotExist = NewError(170000001, "log record not exist")
)
