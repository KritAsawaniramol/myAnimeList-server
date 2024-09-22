package migration

import (
	"log"

	"github.com/kritAsawaniramol/myAnimeList-server/entities"
	"github.com/kritAsawaniramol/myAnimeList-server/pkg/database"
)

func Migration(db database.Database) {
	err := db.GetDb().AutoMigrate(
		&entities.User{},
		&entities.AnimeLists{},
		&entities.Comments{},
		&entities.Credential{},
	)

	if err != nil {
		panic(err)
	}

	log.Println("Database migration completed!")
}
