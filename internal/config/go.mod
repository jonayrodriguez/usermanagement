module github.com/jonayrodriguez/usermanagement/internal/config

go 1.15

require (
	github.com/jonayrodriguez/usermanagement/internal/log v0.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	go.uber.org/zap v1.16.0
	gopkg.in/validator.v2 v2.0.0-20200605151824-2b28d334fa05
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/jonayrodriguez/usermanagement/internal/log => ../log
