package storages

type Storage interface {
	RegisterUser(username, password, email string) error
	AuthenticateUser(username, password string) (bool, error)
}
