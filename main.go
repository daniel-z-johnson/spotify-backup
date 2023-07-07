package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	subLog := log.With().
		Str("Name", "spotify_backup").
		Logger()
	subLog.Info().
		Msg("Application Start")
}
