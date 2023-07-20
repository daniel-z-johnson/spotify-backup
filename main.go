package main

import (
	"github.com/daniel-z-johnson/spotify-backup/controllers"
	"github.com/daniel-z-johnson/spotify-backup/middleware"
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

	router := chi.NewMux()
	example := views.Must(views.ParseFS(templates.TemplateFiles, "main.gohtml", "example.gohtml"))
	staticHandler := http.FileServer(http.FS(templates.StaticFiles))
	router.Use(middleware.Logger(logger))
	router.Get("/example", controllers.StaticPage(example))
	router.Handle("/static/*", staticHandler)
	if err := http.ListenAndServe(":1117", router); err != nil {
		panic(err)
	}
}
