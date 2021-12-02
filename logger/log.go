package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(os.Stdout)

func Log(v ...interface{}) {
	logger.Print(v)
}

func Error(err error) {
	logger.Err(err).Send()
}
