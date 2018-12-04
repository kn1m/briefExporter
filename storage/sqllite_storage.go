package storage

import (
	"briefExporter/configuration"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SqlLiteStorage struct{}

func (s *SqlLiteStorage) SaveUserCredentials(userName *string) error {

	return nil
}

func (s *SqlLiteStorage) SaveUserConfiguration(configuration *UserConfiguration) error {

	return nil
}

func (s *SqlLiteStorage) InitStorage(config *configuration.Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.PathToLocalDb)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db, err
}

func (s *SqlLiteStorage) initScheme(db *sql.DB) error {
	sqlStmt := `
	create table user (id guid not null primary key, name text);
	delete from user;
	`
	_, err := db.Exec(sqlStmt)

	return err
}

func initUserCredentialsTable() error {

	return nil
}

func initUserConfigurationsTable() error {


	return nil
}
