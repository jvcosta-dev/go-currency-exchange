package database

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func Migrate() {
	migrationFiles, err := filepath.Glob("internal/database/migrations/*.sql")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range migrationFiles {
		err := applyMigration(file)
		if err != nil {
			log.Fatalf("err applying migration %s: %v", file, err)
		}
	}
}

func applyMigration(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, err = DB.Exec(string(content))
	return err
}
