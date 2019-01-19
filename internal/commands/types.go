package commands

import "github.com/gtalarico/pm/internal/config"

type Command struct {
	Name     string
	NumArgs  int
	UsageMsg string
	Run      func([]string, config.Config)
}
