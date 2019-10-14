package errors

var (
	NotFoundUserErr       = NewBaseError("用户不存在")
	UserNameOrPasswordErr = NewBaseError("用户不存在或者密码错误")
	AccessTokenErr        = NewBaseError("生成签名错误")
	CreateMemberErr       = NewBaseError("注册失败")
	AccessTokenValidErr                    = NewBaseError("AccessToken 验证失败")
	AccessTokenValidationErrorExpiredErr   = NewBaseError("AccessToken过期")
	AccessTokenValidationErrorMalformedErr = NewBaseError("AccessToken格式错误")
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