package status

import "go-service/common/i18n"

//20x 一定是正确。所有异常都不使用200返回码
//10x 用于指定客户端应相应的某些动作。
//20x 用于表示请求成功。
//30x 表示要完成请求，需要进一步操作（重定向）。
//40x 请求路径问题
//50x 服务器错误。

//错误码的前3位是http status code
const (
	//10x 消息

	//20x 成功----------------------------------------------------------------------------------/
	OK = 0 //成功

	//400xx StatusBadRequest：请求参数有误，当前请求无法被服务器理解
	BadRequest = 400001 // 参数不合法，请检查参数
	//401xx StatusUnauthorized  当前请求需要用户验证
	Unauthorized = 401001 //用户没有登录
	TokenExpired = 401002 //token过期
	NoPermission = 401004 //权限不足

	//403xx StatusForbidden 服务器已经理解请求，但是拒绝执行它
	Forbidden           = 403001 // 通用型, 服务器拒绝执行
	MarketNotDeepEnough = 403002
	VolumeTooLarge      = 403003
	InsufficientBalance = 403004
	InvalidTrade        = 403005
	LockFundError       = 403006
	UnlockFundError     = 403007
	DepositIsExistError = 403008

	StatusNotFound = 404001 //资源不存在

	InternalServerError = 500001
)

var statusTextCN = map[int]string{
	OK: "请求成功",

	BadRequest:          "参数不合法，请检查参数",
	Unauthorized:        "用户没有登录",
	TokenExpired:        "token过期",
	NoPermission:        "权限不足",
	MarketNotDeepEnough: "市场深度不足",
	VolumeTooLarge:      "订单太大",
	InsufficientBalance: "余额不足",
	InvalidTrade:        "无效的交易",
	LockFundError:       "资锁金失败",
	UnlockFundError:     "解锁资金失败",
	DepositIsExistError: "重复的充值请求",
}

var statusTextEN = map[int]string {
	OK: 				"Success",

	BadRequest:          "Bad request",
	Unauthorized:        "Unauthorized",
	TokenExpired:        "token expired",
	NoPermission:        "No permission",
	MarketNotDeepEnough: "Market is not deep enough",
	VolumeTooLarge:      "Volume too large",
	InsufficientBalance: "Insufficient balance",
	InvalidTrade:        "Invalid trade",
	LockFundError:       "Lock fund error",
	UnlockFundError:     "Unlock fund error",
	DepositIsExistError: "deposit is exist error",
}

func StatusText(code, language int) string {
	if language == i18n.CN {
		return statusTextCN[code]
	}
	return statusTextEN[code]
}

func HttpStatueCode(code int) int {
	if OK == code {
		return 200
	}
	return code / 1000
}
