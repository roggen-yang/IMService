package errors

var (
	NotFoundUserErr                        = NewBaseError("用户不存在")
	UserNameOrPasswordErr                  = NewBaseError("用户不存在或者密码错误")
	AccessTokenErr                         = NewBaseError("生成签名错误")
	CreateMemberErr                        = NewBaseError("注册失败")
	AccessTokenValidErr                    = NewBaseError("AccessToken 验证失败")
	AccessTokenValidationErrorExpiredErr   = NewBaseError("AccessToken过期")
	AccessTokenValidationErrorMalformedErr = NewBaseError("AccessToken格式错误")
	UserNoLoginErr                         = NewBaseError("此用户没有登录！")
	SendMessageErr                         = NewBaseError("发送消息失败！")
	PublishMessageErr                      = NewBaseError("发送消息失败")
	UserNotFoundErr                        = NewBaseError("用户不存在")
	ImAddressErr                           = NewBaseError("请配置消息服务地址")
	AddDataErr                             = NewBaseError("维护关系错误")
	imRpcModelMapErr                       = NewBaseError("没有找到对应的RPC服务")
)

type BaseError struct {
	message string
}

func NewBaseError(message string) *BaseError {
	return &BaseError{message: message}
}

func (e *BaseError) Error() string {
	return e.message
}
