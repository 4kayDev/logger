package logger

import (
	"github.com/rs/zerolog"
)
/* envMode = 0..2 
	0 = DebugLevel(Default)
	1 = InfoLevel
	2 = TraceLevel*/
func ConfigureZeroLogger(envMode ...uint8) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	var env uint8
	if len(envMode) > 0 {
		env = envMode[0]
	}
	switch env {
	case 0:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case 1:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case 2:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

