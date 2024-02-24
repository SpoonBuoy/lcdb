package idatabase

type IDatabaseService interface {
	Connect(key string) (string, error)
	PingAll(key string) (string, error)
}

type IDatabaseRepo interface {
	Connect(key string) (string, error)
	PingAll(key string) (string, error)
}
