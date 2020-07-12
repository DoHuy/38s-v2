package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		ListenAddress string `envconfig:"LISTEN_ADDRESS" default:"0.0.0.0:8080"`
		RunMode       string `envconfig:"RUN_MODE" default:"DEV"`
	}

	ViewDir struct {
		Template   string `envconfig:"VIEW_DIR_TEMPLATE" default:"src/38s/template/"`
		Static     string `envconfig:"VIEW_DIR_STATIC" default:"src/38s/static"`
		DomainPath string `envconfig:"VIEW_DIR_DOMAIN_PATH" default:"/"`
	}

	DB struct {
		MySQLConnectionString string `envconfig:"MYSQL_CONNECTION_STRING" required:"false"`
	}

	FCMConfig struct {
		// FCM notification config
		FCMHost        string `envconfig:"FCM_HOST" required:"false"`
	}
	RedisConfig struct {
		DB            int    `envconfig:"REDIS_DB" default:"0"`
		Host          string `envconfig:"REDIS_HOST" default:"192.168.1.8"`
		Port          string `envconfig:"REDIS_PORT" default:"6379"`
		Password      string `envconfig:"REDIS_PASSWORD" default:""`
		SessionSecret string `envconfig:"REDIS_SESSION_SECRET" required:"false"`
		MaxAge        int    `envconfig:"REDIS_SESSION_MAX_AGE" default:"86400"`
	}
}

func GetAuthConfig() (*Config, error) {
	conf := new(Config)
	if err := conf.loadFromEnv(); err != nil {
		return nil, err
	}
	return conf, nil
}

func (c *Config) loadFromEnv() error {
	return envconfig.Process("", c)
}
