package storages

type Storage interface {
	RegisterUser(username, password string) error
	AuthenticateUser(username, password string) (bool, error)
}
