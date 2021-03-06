package dmawatcher

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/trampfox/dma-watcher/models"
)

type WeatherDataStore struct {
	metebridge interface {
		Add(models.MeteoBridge) error
	}
}

func NewWeatherDataStore() (*WeatherDataStore, error) {
	log.Println("Trying to connect to database instance...")
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &WeatherDataStore{
		metebridge: &models.MeteoBridgeModel{ConnPool: pool},
	}, nil
}
