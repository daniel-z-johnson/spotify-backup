package main

import (
	"flag"
	"fmt"
	"github.com/daniel-z-johnson/spotify-backup/templates"
	"go.uber.org/zap"
)

func main() {
	backupFolder := flag.String("location", "~/Documents/spotifyBackup", "location of the backup folder")
	flag.Parse()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Application Start")
	logger.Info("Details", zap.String("folder", *backupFolder))
	t := templates.FS
	dir, _ := t.ReadDir(".")
	fmt.Println(dir[0].Info())
}
