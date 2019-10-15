module github.com/roggen-yang/IMService

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.3.1

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-acme/lego v2.7.2+incompatible // indirect
	github.com/go-xorm/xorm v0.7.9
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/micro/go-micro v1.11.3
	github.com/micro/go-plugins v1.3.0 // indirect
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	gopkg.in/go-playground/validator.v8 v8.18.2
)
