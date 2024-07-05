package main

import (
	"database/sql"
	"flag"
	"github.com/daniel-z-johnson/spotify-backup/conf"
	"github.com/daniel-z-johnson/spotify-backup/models"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/daniel-z-johnson/spotify-backup/controllers"
	"github.com/daniel-z-johnson/spotify-backup/middleware"
	"github.com/daniel-z-johnson/spotify-backup/templates"
	"github.com/daniel-z-johnson/spotify-backup/views"
)

func main() {

	logger := loggerSetup()
	logger.Info("Application Start")

	fileLocation := flag.String("f", "conf.json", "file that contains the configuration")
	flag.Parse()

	config := confSetup(fileLocation)
	db := dbSetup(config)
	sessionRepo := &models.SessionRepo{DB: db}

	router := chi.NewMux()
	example := views.Must(views.ParseFS(templates.TemplateFiles, "main.gohtml", "example.gohtml"))
	oauthPage := views.Must(views.ParseFS(templates.TemplateFiles, "main.gohtml", "oauthpage.gohtml"))
	staticHandler := http.FileServer(http.FS(templates.StaticFiles))
	music := &controllers.Music{}
	music.Templates.OAuthPage = oauthPage
	sessionStore := middleware.SessionStore{Session: sessionRepo}

	router.Use(middleware.Logger(logger))
	router.Use(middleware.Session())
	router.Use(sessionStore.HandleState())

	router.Get("/example", controllers.StaticPage(example))
	router.Get("/oauth", music.OauthPage)
	router.Handle("/static/*", staticHandler)
	if err := http.ListenAndServe(":1117", router); err != nil {
		panic(err)
	}
}

func loggerSetup() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger = logger.With(slog.Any("application", "spotify_backup"))
	return logger
}

func confSetup(fileLocation *string) *conf.Conf {
	config, err := conf.LoadConf(fileLocation)
	if err != nil {
		// nothing can be done just give for now
		panic(err)
	}
	return config
}

func dbSetup(config *conf.Conf) *sql.DB {
	db, err := models.Open(config)
	if err != nil {
		panic(err)
	}
	return db
}
