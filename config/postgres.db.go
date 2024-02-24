package config

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func NewDBConfig(host string, user string, pwd string, name string, port string) *DBConfig {
	return &DBConfig{
		Host:     host,
		User:     user,
		Password: pwd,
		Name:     name,
		Port:     port,
	}
}
