package models

type Repo struct {
	Id             uint64 `gorm:"repo_id;primary_key"`
	Name           string `gorm:"name"`
	Link           string `gorm:"link"`
	CloneHttpUrl   string `gorm:"clone_http_url"`
	UserId         uint64 `gorm:"user_id"`
	LatestCommitId string `gorm:"latest_commit_id"`
	LineCount      uint64 `gorm:"line_count"`
}

type Commit struct {
	Id           uint64   `gorm:"commit_id;primary_key"`
	CommitId     string   `gorm:"commit_id"`
	UserId       string   `gorm:"user_id"`
	RepoId       string   `gorm:"repo_id"`
	LinesAdded   uint64   `gorm:"lines_added"`
	LinesRemoved uint64   `gorm:"lines_removed"`
	LineCount    uint64   `gorm:"line_count"`
	Languages    []string `gorm:"languages"`
}
