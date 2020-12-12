package models

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MeteoBridgeModel struct {
	ConnPool *pgxpool.Pool
}

// MeteoBridge model holds data from meteobridge weather stations
type MeteoBridge struct {
	Info map[string]interface{} // JSONB field
}

func (m *MeteoBridgeModel) All() ([]MeteoBridge, error) {
	// TODO add timeout
	rows, err := m.ConnPool.Query(context.Background(), "SELECT * FROM meteobridgedata")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []MeteoBridge

	for rows.Next() {
		var item MeteoBridge

		if err = rows.Scan(&item.Info); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (m *MeteoBridgeModel) Add(item MeteoBridge) error {
	infoJson, err := json.Marshal(item.Info)
	if err != nil {
		return err
	}

	log.Println("Creating a new record in the meteobridgedata table...")
	_, err = m.ConnPool.Exec(context.Background(),
		"INSERT INTO meteobridgedata (info) VALUES ($1)", &infoJson)
	if err != nil {
		return err
	}

	return nil
}
