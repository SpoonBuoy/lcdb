package service

type Auth interface{
	SetToken() error
	GetToken() error

}


type GithubAuth struct {}
func (gta *GithubAuth) SetGithubToken() error {
	return nil
}

func (gta *GithubAuth) GetGithubToken() error {
	return nil
}
