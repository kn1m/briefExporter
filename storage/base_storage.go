package storage

type User struct {
	Email string
}

type UserConfiguration struct {
	ScanFolder string
}

type Storage interface {
	SaveUserCredentials(userName *string) error
	SaveUserConfiguration(configuration *UserConfiguration) error
}