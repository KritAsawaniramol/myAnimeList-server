package main

import (
	"log"
	"os"

	"github.com/kritAsawaniramol/myAnimeList-server/config"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/database"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/database/migration"
	"github.com/kritAsawaniramol/myAnimeList-server/server"
)

func main() {
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	db := database.NewPostgresDatabase(&cfg)

	migration.Migration(db)

	server.NewGinServer(&cfg, db.GetDb()).Start()
}
