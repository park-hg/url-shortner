package pkg

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
	mysqlgorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"traffic-reporter/config"
)

func MustConnectMySQL(c config.MySQLConfig) *gorm.DB {
	sqlDB, err := sql.Open("mysql", c.DataSourceName)

	gDB, err := gorm.Open(
		mysqlgorm.New(mysqlgorm.Config{Conn: sqlDB}),
		&gorm.Config{
			Logger:                 setLoggerFromOption(c.LoggingLevel),
			SkipDefaultTransaction: true,
			TranslateError:         true,
		},
	)
	if err != nil {
		panic(err)
	}

	return gDB
}

func setLoggerFromOption(level string) logger.Interface {
	loggerLevel := logger.Silent
	switch level {
	case "debug", "info":
		loggerLevel = logger.Info
	case "warn":
		loggerLevel = logger.Warn
	case "error":
		loggerLevel = logger.Error
	}

	if loggerLevel == logger.Silent {
		return logger.Discard
	}
	return logger.Default.LogMode(loggerLevel)
}

func MustConnectRedis(c config.RedisConfig) redis.UniversalClient {
	redisClient := redis.NewClient(&redis.Options{Addr: c.Endpoint, ClientName: c.ClientName})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return redisClient
}
