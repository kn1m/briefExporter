package storage

type User struct {
	Email string
}

type UserConfiguration struct {
}

type Storage interface {
	SaveUserCredentials() error
	SaveUserConfiguration(configuration *UserConfiguration) error
}
