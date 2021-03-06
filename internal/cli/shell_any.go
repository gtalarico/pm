// +build !windows

package cli

import (
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"syscall"

	"github.com/pkg/errors"
)

func handleShellError() {
	shellError := recover()
	if shellError != nil {
		err := errors.New("shell error")
		Abort(err)
	}
}

func silentCtrlC() {
	// listen to terminate signales
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// and don't so anything
	go func() { <-c }()
}

func Shell(cwd string) {
	//technosophos.com/2014/07/11/start-an-interactive-shell-from-within-go.html
	defer handleShellError()

	// silent the ctrl+c / SIGTERM signals
	silentCtrlC()

	fmt.Println("Starting new shell")
	fmt.Println("Use 'exit' to terminate child shell")

	// Get the current user.
	me, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Set an environment variable.
	// os.Setenv("SOME_VAR", "1")

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	// Start up a new shell.
	// Note that we supply "login" twice.
	// -fpl means "don't prompt for PW and pass through environment."
	// fmt.Print(">> Starting a new interactive shell")
	proc, err := os.StartProcess("/usr/bin/login", []string{"login", "-fpl", me.Username}, &pa)
	if err != nil {
		panic(err)
	}

	// Wait until user exits the shell
	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	// os.Setenv("PROMPT", "()")
	// Keep on keepin' on.
	fmt.Printf("Exited Go Sub Shell\n %s\n", state.String())
}
