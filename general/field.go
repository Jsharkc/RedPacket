package general

const (
	IntZero = 0
	IntOne  = 1

	// TokenExpireInHour - token active time (hour)
	TokenExpireInHour = 24

	// Respone key
	RespTokenKey = "token"
	RespBalance  = "balance"
	RespRPID     = "rpid"
	RespPwd      = "pwd"
	RespMoney    = "money"
	RespGrabList = "grablist"

	ClaimUID    = "uid"
	ClaimExpire = "exp"

	Dialect      = "mysql"
	MysqlArg     = "root:ziyao945@tcp(127.0.0.1:3306)/redpack?charset=utf8&parseTime=true"
	MysqlTestArg = "root:ziyao945@tcp(127.0.0.1:3306)/redpack?charset=utf8&parseTime=true"

	// RedPack status
	RPGrab   = 0x20
	RPRefund = 0x21
	RPFinish = 0x22

	// Red packet active time
	RPActiveTimeInHour = 24

	LetterSource = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	TargetPWDNum = 8

	// mysql table column name
	TBTotalMoney        = "totalmoney"
	TBTotalMoneyExprSub = "totalmoney - ?"
	TBNumber            = "num"
	TBNumberExprSub     = "num - ?"
	TBStatus            = "status"
	TBBalance           = "balance"
	TBBalanceExprAdd    = "balance + ?"

	// Timer work param
	TimerWorkErrorTryTime     = 5
	TimerWorkErrorTryInterval = 200
)
