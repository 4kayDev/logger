package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime, }).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Int("pid", os.Getpid()).
	Logger()

func Warn(msg string, tag ...string) {
	if len(tag) == 4 {
		Logger.Warn().Str(tag[0], tag[1]).Str(tag[2],tag[3]).Msg(msg)
		return
	}
	if len(tag) == 3 {
		Logger.Warn().Str(tag[0], tag[1]).Str("SRC",tag[2]).Msg(msg)
		return
	}
	if len(tag) == 2 {
		Logger.Warn().Str(tag[0], tag[1]).Msg(msg)
		return
	}
	if len(tag) == 1 {
		Logger.Warn().Str("SRC", tag[0]).Msg(msg)
		return
	}
	Logger.Warn().Msg(msg)
}

func Info(msg string, tag ...string) {
	if len(tag) == 4 {
		Logger.Info().Str(tag[0], tag[1]).Str(tag[2],tag[3]).Msg(msg)
		return
	}
	if len(tag) == 3 {
		Logger.Info().Str(tag[0], tag[1]).Str("SRC",tag[2]).Msg(msg)
		return
	}
	if len(tag) == 2 {
		Logger.Info().Str(tag[0], tag[1]).Msg(msg)
		return
	}
	if len(tag) == 1 {
		Logger.Info().Str("SRC", tag[0]).Msg(msg)
		return
	}
	Logger.Info().Msg(msg)
}

func Debug(msg string, tag ...string) {
	if len(tag) == 4 {
		Logger.Debug().Str(tag[0], tag[1]).Str(tag[2],tag[3]).Msg(msg)
		return
	}
	if len(tag) == 3 {
		Logger.Debug().Str(tag[0], tag[1]).Str("SRC",tag[2]).Msg(msg)
		return
	}
	if len(tag) == 2 {
		Logger.Debug().Str(tag[0], tag[1]).Msg(msg)
		return
	}
	if len(tag) == 1 {
		Logger.Debug().Str("SRC", tag[0]).Msg(msg)
		return
	}
	Logger.Debug().Msg(msg)
}

func Trace(msg string, tag ...string) {
	if len(tag) == 4 {
		Logger.Trace().Str(tag[0], tag[1]).Str(tag[2],tag[3]).Msg(msg)
		return
	}
	if len(tag) == 3 {
		Logger.Trace().Str(tag[0], tag[1]).Str("SRC",tag[2]).Msg(msg)
		return
	}
	if len(tag) == 2 {
		Logger.Trace().Str(tag[0], tag[1]).Msg(msg)
		return
	}
	if len(tag) == 1 {
		Logger.Trace().Str("SRC", tag[0]).Msg(msg)
		return
	}
	Logger.Trace().Msg(msg)
}

func Panic(err error, msg string, tag ...string) {
	if len(tag) == 4 {
		Logger.Panic().Err(err).Str(tag[0], tag[1]).Str(tag[2],tag[3]).Msg(msg)
		return
	}
	if len(tag) == 3 {
		Logger.Panic().Err(err).Str(tag[0], tag[1]).Str("SRC",tag[2]).Msg(msg)
		return
	}
	if len(tag) == 2 {
		Logger.Panic().Err(err).Str(tag[0], tag[1]).Msg(msg)
		return
	}
	if len(tag) == 1 {
		Logger.Panic().Err(err).Str("SRC", tag[0]).Msg(msg)
		return
	}
	Logger.Panic().Err(err).Msg(msg)
}

func Fatal(err error, msg string, tag ...string) {
	if len(tag) == 4 {
		Logger.Fatal().Err(err).Str(tag[0], tag[1]).Str(tag[2],tag[3]).Msg(msg)
		return
	}
	if len(tag) == 3 {
		Logger.Fatal().Err(err).Str(tag[0], tag[1]).Str("SRC",tag[2]).Msg(msg)
		return
	}
	if len(tag) == 2 {
		Logger.Fatal().Err(err).Str(tag[0], tag[1]).Msg(msg)
		return
	}
	if len(tag) == 1 {
		Logger.Fatal().Err(err).Str("SRC", tag[0]).Msg(msg)
		return
	}
	Logger.Fatal().Err(err).Msg(msg)
}

func Error(err error, msg string, tag ...string) {
	if len(tag) == 4 {
		Logger.Error().Err(err).Str(tag[0], tag[1]).Str(tag[2],tag[3]).Msg(msg)
		return
	}
	if len(tag) == 3 {
		Logger.Error().Err(err).Str(tag[0], tag[1]).Str("SRC",tag[2]).Msg(msg)
		return
	}
	if len(tag) == 2 {
		Logger.Error().Err(err).Str(tag[0], tag[1]).Msg(msg)
		return
	}
	if len(tag) == 1 {
		Logger.Error().Err(err).Str("SRC", tag[0]).Msg(msg)
		return
	}
	Logger.Error().Err(err).Msg(msg)
}
