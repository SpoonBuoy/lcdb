package models

type GithubMetric struct {
	Id         uint64 `gorm:"mid;primary_key"`
	TotalLines uint64 `gorm:"total_lines"`
}
