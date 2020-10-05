package config

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

const (
	defaultServerPort = 8080
	defaultServerHost = "127.0.0.1"
)

// Config represents an application configuration.
type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

// Server represents an server configuration.
type Server struct {
	Port int    `yaml:"port" envconfig:"SERVER_PORT" validate:"nonzero"`
	Host string `yaml:"host" envconfig:"SERVER_HOST" validate:"nonzero"`
}

// Database represents an database configuration.
type Database struct {
	Username string `yaml:"user" envconfig:"DB_USERNAME" validate:"nonzero"`
	Password string `yaml:"pass" envconfig:"DB_PASSWORD" validate:"nonzero"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validator.Validate(&c)

}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger *logrus.Logger) (*Config, error) {
	// default config
	c := Config{
		Server: Server{
			Port: defaultServerPort,
			Host: defaultServerHost,
		},
		Database: Database{
			Username: "user",
			Password: "password",
		},
	}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Fatal(err.Error())
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		logger.Fatal(err.Error())

		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = envconfig.Process("", &c); err != nil {
		logger.Fatal(err.Error())
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		logger.Fatal(err.Error())
		return nil, err
	}

	return &c, err
}
