package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := log.With().
		Str("Name", "spotify_backup").
		Logger()
	logger.Info().
		Msg("Application Start")
}
