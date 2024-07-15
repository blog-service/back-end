package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var log zerolog.Logger

func ConsoleLog() *zerolog.Logger {
	log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &log
}
