package models

type File struct {
	Path    string
	Content string
	Objects []string
}

type Repo struct {
	Files []File
}

type RepoCredentials struct {
	Url       string
	Owner     string
	Name      string
	CommitSHA string
}

type Data struct {
	Url          string
	InputLines   []string
	Credentials  []RepoCredentials
	Repositories []Repo
	Output       string
}
