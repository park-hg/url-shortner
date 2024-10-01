package config

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	ServerConfig ServerConfig
	FiberConfig  fiber.Config
	MySQLConfig  MySQLConfig
}

type ServerConfig struct {
	Host string
}

type MySQLConfig struct {
	DataSourceName string
	LoggingLevel   string
}

func NewConfig() Config {
	return Config{}
}
