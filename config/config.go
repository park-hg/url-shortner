package config

import (
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig ServerConfig `envPrefix:"SERVER_"`
	FiberConfig  fiber.Config
	MySQLConfig  MySQLConfig `envPrefix:"MYSQL_"`
	RedisConfig  RedisConfig `envPrefix:"REDIS_"`
}

type ServerConfig struct {
	Host string `env:"HOST"`
}

type MySQLConfig struct {
	User         string `env:"USER" envDefault:"root"`
	Password     string `env:"PASSWORD"`
	Host         string `env:"HOST"`
	Port         int    `env:"PORT"`
	Database     string `env:"DATABASE"`
	LoggingLevel string `env:"LOGGING_LEVEL"`
}

type RedisConfig struct {
	Endpoint   string `env:"ENDPOINT"`
	ClientName string `env:"CLIENT_NAME"`
}

func MustNewConfig() Config {
	path := filepath.Dir(packagePath()) // parent directory of the package, which is the project root
	path = filepath.Join(path, ".env")
	if err := godotenv.Load(path); err != nil {
		panic(err)
	}

	var cfg Config
	cfg.FiberConfig = fiber.Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return cfg
}

func packagePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
