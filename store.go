package dmawatcher

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/trampfox/dma-watcher/models"
)

type WeatherDataStore struct {
	metebridge interface {
		Add(models.MeteoBridge) (error)
	}
}

func NewWeatherDataStore() (*WeatherDataStore, error) {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &WeatherDataStore{
		metebridge: &models.MeteoBridgeModel{DB: db},
	}, nil
}
