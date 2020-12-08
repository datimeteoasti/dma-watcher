package models

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type MeteoBridgeModel struct {
	DB *pgx.Conn
}

// MeteoBridge model holds data from meteobridge weather stations
type MeteoBridge struct {
	Info map[string]interface{}	// JSONB field
}

func (m *MeteoBridgeModel) All() ([]MeteoBridge, error) {
	// TODO add timeout
	rows, err := m.DB.Query(context.Background(), "SELECT * FROM meteobridgedata")
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
	// TODO add timeout
	infoJson, err := json.Marshal(item.Info)
	_, err = m.DB.Exec(context.Background(),
		"INSERT INTO meteobridgedata (foo) VALUES ('%s')", infoJson)
	if err != nil {
		return err
	}

	return nil
}
