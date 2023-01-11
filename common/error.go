package common

// Error 错误结构体
type Error struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

// Error error interface
func (e Error) Error() string {
	return e.Msg
}

// 错误常量
const (
	// HttpStatusSuccess http status 返回值 成功
	HttpStatusSuccess = 200
	// HttpStatusInputError http status 返回值 客户端错误
	HttpStatusInputError = 400
	// HttpStatusServerError http status 返回值 服务器错误
	HttpStatusServerError = 500
)

const (
	// DefaultSuccessCode 默认成功编码
	DefaultSuccessCode = 0
	// DefaultErrorCode 默认失败编码
	DefaultErrorCode = iota + 100000
	// ParamsErrorCode 参数错误
	ParamsErrorCode
	// RPCServerNotExistErrorCode RPC配置不存在
	RPCServerNotExistErrorCode
	// GetRPCClientErrorCode RPC客户端获取失败
	GetRPCClientErrorCode
	// RemoteCallErrorCode 下游调用失败
	RemoteCallErrorCode
	// FreqControlErrorCode 频控
	FreqControlErrorCode
	// SignParamErrorCode 签名参数错误
	SignParamErrorCode
	// SignTimeOutErrorCode 签名已过期
	SignTimeOutErrorCode
	// SignChannleErrorCode 签名渠道方错误
	SignChannleErrorCode
	// SignError 签名错误
	SignErrorCode
	// UserNotFoundErrorCode 用户不存在
	UserNotFoundErrorCode
	// UserUpdateErrorCode 用户更新失败
	UserUpdateErrorCode
)

// 自定义错误信息
var (
	Success                = Error{HttpStatusSuccess, DefaultSuccessCode, "success"}
	ParamsError            = Error{HttpStatusInputError, ParamsErrorCode, "params error"}
	ServerError            = Error{HttpStatusServerError, DefaultErrorCode, "service internal err"}
	RPCServerNotExistError = Error{HttpStatusServerError, RPCServerNotExistErrorCode, "RPC server no exist"}
	GetRPCClientError      = Error{HttpStatusServerError, GetRPCClientErrorCode, "get RPC client error"}
	RemoteCallError        = Error{HttpStatusServerError, RemoteCallErrorCode, "remote call error"}
	FreqControlError       = Error{HttpStatusServerError, FreqControlErrorCode, "freq control error"}
	SignParamError         = Error{HttpStatusServerError, SignParamErrorCode, "sign param error"}
	SignTimeOutError       = Error{HttpStatusServerError, SignTimeOutErrorCode, "sign time expire"}
	SignChannleError       = Error{HttpStatusServerError, SignChannleErrorCode, "sign channel error"}
	SignError              = Error{HttpStatusServerError, SignErrorCode, "sign error"}
	UserNotFoundError      = Error{HttpStatusServerError, UserNotFoundErrorCode, "user not found error"}
	UserUpdateError        = Error{HttpStatusServerError, UserUpdateErrorCode, "user update error"}
)

// 可显示给请求方的错误编码
var DisplayableErrorCodes = map[int]interface{}{
	DefaultSuccessCode:    nil,
	DefaultErrorCode:      nil,
	ParamsErrorCode:       nil,
	FreqControlErrorCode:  nil,
	SignParamErrorCode:    nil,
	SignTimeOutErrorCode:  nil,
	SignChannleErrorCode:  nil,
	SignErrorCode:         nil,
	UserNotFoundErrorCode: nil,
	UserUpdateErrorCode:   nil,
}

// FormatError 过滤掉不展示给外界的错误信息，格式化成统一的错误信息输出
func FormatError(err interface{}) error {
	result, ok := err.(*Error)
	if !ok {
		return &ServerError
	}

	// 过滤掉不可展示给用户的错误编码
	if _, ok := DisplayableErrorCodes[result.Code]; !ok {
		return &ServerError
	}
	return result
}
