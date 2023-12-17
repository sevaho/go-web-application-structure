package logger

import (
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/sevaho/gowas/src/config"
)

var Logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "levelname"
	zerolog.TimestampFieldName = "asctime"

    // If running in production, use JSON format as output for logs otherwise pretty print with colors for readability
	if config.Config.ENVIRONMENT != "PRODUCTION" {
		Logger = zerolog.New(io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr})).With().Timestamp().Logger()
	} else {
		zerolog.LevelFieldMarshalFunc = func(l zerolog.Level) string {
			return strings.ToUpper(l.String())
		}
		Logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	}
}
