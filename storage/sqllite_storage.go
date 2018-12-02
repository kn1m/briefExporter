package storage

import (
	_ "github.com/mattn/go-sqlite3"
)

type SqlLiteStorage struct{}

func (s *SqlLiteStorage) SaveUserCredentials() error {

	return nil
}

func (s *SqlLiteStorage) SaveUserConfiguration(configuration *UserConfiguration) error {

	return nil
}
