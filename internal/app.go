package internal

import (
	"traffic-reporter/config"
	"traffic-reporter/internal/pkg"
	"traffic-reporter/internal/shortener/adapter"
	"traffic-reporter/internal/shortener/usecase"
)

type App struct {
	config config.Config

	ShortenURLUseCase *usecase.ShortenURLUseCase
}

func InitApp(c config.Config) *App {
	db := pkg.MustConnectMySQL(c.MySQLConfig)
	rdb := pkg.MustConnectRedis(c.RedisConfig)
	idGenerator := pkg.NewTSIDGenerator(rdb, pkg.NewBase62Encoder())
	repo := adapter.NewMySQLURLRepository(db, idGenerator)
	shortenURLUseCase := usecase.NewShortenURLUseCase(repo)
	return &App{
		config: c,

		ShortenURLUseCase: shortenURLUseCase,
	}
}
