package storage

import (
	"context"
	"encoding/json"
	"main/core"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	Pool *pgxpool.Pool
}

func (s *Store) SaveNewRecord(r *core.Record) error {
	info, err := json.Marshal(r.Info)
	if err != nil {
		return err
	}
	_, err = s.Pool.Exec(context.Background(), "INSERT INTO records (date, info) VALUES ($1, $2)", r.Date, info)
	if err != nil {
		return err
	}
	return nil
}
