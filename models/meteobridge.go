package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type MeteoBridgeModel struct {
	DB *pgx.Conn
}

// MeteoBridge model holds data from meteobridge weather stations
type MeteoBridge struct {
	Foo string
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

		if err = rows.Scan(&item.Foo); err != nil {
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
	_, err := m.DB.Query(context.Background(), 
		fmt.Sprintf("INSERT INTO meteobridgedata (foo) VALUES ('%s')", item.Foo))
	if err != nil {
		return err
	}

	return nil
}