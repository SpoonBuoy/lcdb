package models

type User struct {
	Id             uint64 `gorm:"user_id;primary_key"`
	Name           string `gorm:"name"`
	Password       string `gorm:"password"`
	Email          string `gorm:"email"`
	GithubUserName string `gorm:"github_user_name"`
}
