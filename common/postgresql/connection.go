package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DBname,
		config.MaxConnection,
		config.MaxConnectionIdleTime)

	connConfig, parseConfigError := pgxpool.ParseConfig(connString) // checks if connString has problems
	if parseConfigError != nil {
		panic(parseConfigError)
	}

	conn, err := pgxpool.ConnectConfig(context, connConfig) // actually makes the connection.
	if err != nil {
		log.Error("unable to connect db: %v\n", err)
		panic(err)
	}

	return conn
}
