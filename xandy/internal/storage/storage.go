package storage

import (
	"context"

	"github.com/eac0de/xandy/shared/pkg/psql"
)

type xandyStorage struct {
	*psql.PSQLStorage
}

func NewxandyStorage(
	ctx context.Context,
	host string,
	port string,
	username string,
	password string,
	dbName string,
) (*xandyStorage, error) {
	storage, err := psql.New(ctx, host, port, username, password, dbName)
	if err != nil {
		return nil, err
	}
	return &xandyStorage{storage}, nil
}
