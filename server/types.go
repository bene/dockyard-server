package server

type Project struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	GitRepo  string `json:"git_repo"`
	IsPublic bool   `json:"is_public"`
}

type BuildStatus string

const (
	FAILED   = BuildStatus("FAILED")
	PENDING  = BuildStatus("PENDING")
	FINISHED = BuildStatus("FINISHED")
)

type Build struct {
	Id        string      `json:"id"`
	ProjectId string      `json:"project_id"`
	Tags      *[]string   `json:"tags"`
	Status    BuildStatus `json:"status"`
	Date      int64       `json:"date"`
}
