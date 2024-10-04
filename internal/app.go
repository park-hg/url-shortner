package internal

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"traffic-reporter/config"
	"traffic-reporter/internal/pkg"
	"traffic-reporter/internal/shortener/adapter"
	"traffic-reporter/internal/shortener/usecase"
)

type App struct {
	config config.Config

	db          *gorm.DB
	rdb         redis.UniversalClient
	idGenerator pkg.IDGenerator

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

		db:          db,
		rdb:         rdb,
		idGenerator: idGenerator,

		ShortenURLUseCase: shortenURLUseCase,
	}
}

func (a *App) Teardown() error {
	if err := a.idGenerator.Close(); err != nil {
		return err
	}

	sqlDB, err := a.db.DB()
	if err != nil {
		return err
	}
	if err = sqlDB.Close(); err != nil {
		return err
	}

	if err = a.rdb.Close(); err != nil {
		return err
	}

	return nil
}
