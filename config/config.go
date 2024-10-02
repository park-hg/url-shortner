package config

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	ServerConfig ServerConfig
	FiberConfig  fiber.Config
	MySQLConfig  MySQLConfig
	RedisConfig  RedisConfig
}

type ServerConfig struct {
	Host string
}

type MySQLConfig struct {
	DataSourceName string
	LoggingLevel   string
}

type RedisConfig struct {
	Endpoint   string
	ClientName string
}

func NewConfig() Config {
	return Config{}
}
