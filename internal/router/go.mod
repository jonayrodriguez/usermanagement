module github.com/jonayrodriguez/usermanagement/internal/router

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jonayrodriguez/usermanagement/internal/controller v0.0.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.1.13 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect

)

replace (
	github.com/jonayrodriguez/usermanagement/internal/config => ../config
	github.com/jonayrodriguez/usermanagement/internal/controller => ../controller
	github.com/jonayrodriguez/usermanagement/pkg/model => ../../pkg/model
	github.com/jonayrodriguez/usermanagement/internal/database => ../../internal/database

)
