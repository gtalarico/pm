package cli

import "github.com/pkg/errors"

func ValidateArgs(args []string) (cmdName string, posArgs []string, err error) {

	if len(args) == 0 {
		err = errors.New("missing command name")
		return
	}

	if len(args) > 1 {
		posArgs = args[1:]
	}
	cmdName = args[0]
	return

}
