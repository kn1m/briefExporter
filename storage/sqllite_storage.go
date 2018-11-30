package storage

type SqlLiteStorage struct{}

func Init() {

}

func (s *SqlLiteStorage) SaveUserCredentials() error {

	return nil
}

func (s *SqlLiteStorage) SaveUserConfiguration(configuration *UserConfiguration) error {

	return nil
}
