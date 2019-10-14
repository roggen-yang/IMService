package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/roggen-yang/IMService/common/constant"
	"github.com/roggen-yang/IMService/common/errors"
	"github.com/roggen-yang/IMService/common/response"
)

func ValidAccessToken(context *gin.Context) {
	authorization := context.GetHeader(constant.DefaultField)
	token, err := jwt.Parse(authorization, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(constant.UserSignedKey), nil
	})
	if err != nil {
		if err, ok := err.(*jwt.ValidationError); ok {
			if err.Errors&jwt.ValidationErrorMalformed != 0 {
				response.HttpResponse(context, nil, errors.AccessTokenValidationErrorMalformedErr)
				context.Abort()
				return
			}
			if err.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				response.HttpResponse(context, nil, errors.AccessTokenValidationErrorExpiredErr)
				context.Abort()
				return
			}
		}
		response.HttpResponse(context, nil, errors.AccessTokenValidErr)
		context.Abort()
		return
	}
	if token != nil && token.Valid {
		context.Next()
		return
	}
	response.HttpResponse(context, nil, errors.AccessTokenValidErr)
	context.Abort()
	return

}
