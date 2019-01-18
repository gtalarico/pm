package cmd

type Config struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Path string `json:"path"`
}
