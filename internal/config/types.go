package config

type Config struct {
	Projects []*Project `json:"projects"`
}

type Project struct {
	Path string `json:"path"`
	Name string `json:"name"`
	// Command string `json:"command"`
}
