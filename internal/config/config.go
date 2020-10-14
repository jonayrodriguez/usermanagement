package config

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

// Config represents an application configuration.
type Config struct {
	App      App      `yaml:"app"`
	Server   Server   `yaml:"server"`
	Logging  Logging  `yaml:"logging"`
	Database Database `yaml:"database"`
}

type App struct {
	Name       string `yaml:"name" envconfig:"APP_NAME" validate:"nonzero"`
	Version    string `yaml:"version" envconfig:"APP_VERSION" validate:"nonzero"`
	Repository string `yaml:"repository" envconfig:"APP_REPOSITORY" validate:"nonzero"`
}

// Server represents an server configuration.
type Server struct {
	Port        int    `yaml:"port" envconfig:"SERVER_PORT" validate:"nonzero"`
	Host        string `yaml:"host" envconfig:"SERVER_HOST" validate:"nonzero"`
	Environment string `yaml:"environment" envconfig:"SERVER_ENV" validate:"nonzero"`
}

// Logging represents an logging configuration.
type Logging struct {
	File         string `yaml:"file" envconfig:"LOGGING_FILE" validate:"nonzero"`
	Format       string `yaml:"format" envconfig:"LOGGING_FORMAT" validate:"nonzero"`
	Level        string `yaml:"level" envconfig:"LOGGING_LEVEL" validate:"nonzero"`
	AccessFile   string `yaml:"accessFile" envconfig:"LOGGING_ACCESS_FILE" validate:"nonzero"`
	AccessFormat string `yaml:"accessFormat" envconfig:"LOGGING_ACCESS_FORMAT" validate:"nonzero"`
	AccessLevel  string `yaml:"accessLevel" envconfig:"LOGGING_ACCESS_LEVEL" validate:"nonzero"`
	MaxSize      int    `yaml:"maxSize" envconfig:"LOGGING_MAX_SIZE" validate:"nonzero"`
	MaxBackup    int    `yaml:"maxBackup" envconfig:"LOGGING_MAX_BACKUP" validate:"nonzero"`
	MaxAge       int    `yaml:"maxAge" envconfig:"LOGGING_MAX_AGE" validate:"nonzero"`
}

// Database represents an database configuration.
type Database struct {
	Username string `yaml:"username" envconfig:"DB_USERNAME" validate:"nonzero"`
	Password string `yaml:"password" envconfig:"DB_PASSWORD" validate:"nonzero"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validator.Validate(&c)

}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string) (*Config, error) {
	// default empty config
	c := Config{}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = envconfig.Process("", &c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}
