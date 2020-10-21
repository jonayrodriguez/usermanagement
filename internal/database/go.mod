module github.com/jonayrodriguez/usermanagement/internal/database

go 1.15

require (

	github.com/jonayrodriguez/usermanagement/pkg/model v0.0.0
	github.com/jonayrodriguez/usermanagement/internal/config v0.0.0

	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	gopkg.in/validator.v2 v2.0.0-20200605151824-2b28d334fa05 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/mysql v1.0.2
	gorm.io/gorm v1.20.3


)

replace (
	github.com/jonayrodriguez/usermanagement/pkg/model => ../../pkg/model

	github.com/jonayrodriguez/usermanagement/internal/config => ../config

)