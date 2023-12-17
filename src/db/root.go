package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/avast/retry-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sevaho/gowas/src/config"
	"github.com/sevaho/gowas/src/logger"
)

func SetUpDB() *Queries {
	conn, err := pgxpool.New(context.Background(), config.Config.DB_DSN)

	if err != nil {
		panic(err)
	}

	db := New(conn)

	err = retry.Do(
		func() (err error) {
			ctx := context.Background()
			_, err = db.db.Exec(ctx, "SELECT NOW()")
			if err != nil {
				logger.Logger.Warn().Err(err).Msgf("Unable to connect to db %s", strings.Split(config.Config.DB_DSN, "@")[1])
			}
			return
		},
		retry.Attempts(uint(config.Config.DB_CONNECTION_RETRIES)),
		retry.Delay(config.Config.DB_CONNECTION_RETRY_DELAY),
	)
	if err != nil {
		msg := fmt.Sprintf("Unable to connect to db %s", strings.Split(config.Config.DB_DSN, "@")[1])
		panic(msg)
	}

	logger.Logger.Info().Msgf("[DB] connected to %s", strings.Split(config.Config.DB_DSN, "@")[1])

	return db
}
