package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectToDb(ctx context.Context) (*pgx.Conn, error) {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db %v", err)
	}
	defer conn.Close(context.Background())

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to ping db %v", err)
	}

	return conn, nil
}
