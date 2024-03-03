package db

type Store interface {
	Connect() (Store, error)
	Close() error
	GetClient() error
}
