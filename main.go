package main

import (
	"github.com/daniel-z-johnson/spotify-backup/templates"
	"github.com/daniel-z-johnson/spotify-backup/views"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := log.With().
		Str("Name", "spotify_backup").
		Logger()
	logger.Info().
		Msg("Application Start")

	handler := chi.NewMux()
	example, err := views.ParseFS(templates.TemplateFiles, "main.gohtml", "example.gohtml")
	if err != nil {
		panic("Template issue: " + err.Error())
	}
	handler.HandleFunc("/example", func(resp http.ResponseWriter, req *http.Request) {
		example.Execute(resp, nil)
	})
	if err := http.ListenAndServe(":1117", handler); err != nil {
		panic(err)
	}
}
