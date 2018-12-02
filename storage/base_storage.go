package storage

import (
	"briefExporter/configuration"
	"database/sql"
	"log"
)

type User struct {
	Email string
}

type UserConfiguration struct {
}

type Storage interface {
	SaveUserCredentials() error
	SaveUserConfiguration(configuration *UserConfiguration) error
}

func InitStorage(config *configuration.Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.PathToLocalDb)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db, err
}

func initScheme(db *sql.DB) error {
	sqlStmt := `
	create table user (id guid not null primary key, name text);
	delete from user;
	`
	_, err := db.Exec(sqlStmt)

	return err
}
