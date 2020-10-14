module github.com/jonayrodriguez/usermanagement

go 1.15

require (
	github.com/gin-contrib/zap v0.0.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jonayrodriguez/usermanagement/internal/config v0.0.0
	github.com/jonayrodriguez/usermanagement/internal/log v0.0.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/ugorji/go v1.1.11 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee // indirect
	golang.org/x/sys v0.0.0-20201013132646-2da7054afaeb // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0 // indirect

)

replace github.com/jonayrodriguez/usermanagement/internal/config => ../internal/config

replace github.com/jonayrodriguez/usermanagement/internal/log => ../internal/log
