package cli

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/pkg/errors"
)

func handleShellError() {
	shellError := recover()
	if shellError != nil {
		err := errors.New("shell error")
		Abort(err)
	}
}

func Shell(cwd string) {
	//technosophos.com/2014/07/11/start-an-interactive-shell-from-within-go.html
	defer handleShellError()

	fmt.Println("Starting new shell")
	fmt.Println("Use 'CTRL + C' or '$ exit' to terminate child shell")

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

	stdin := os.Stdin
	stdout := os.Stdout
	defer func() {
		os.Stdin = stdin
		os.Stdout = stdout
	}()

	r, w, _ := os.Pipe()
	os.Stdin = r
	os.Stdout = w

	time.Sleep(1 * time.Second)
	// var in string
	// in = "ls"
	// w.WriteString("ls\n\n")
	w.Write([]byte("ls\n"))
	// w.Close()
	// fmt.Scan(&in)
	// fmt.Println("ls")
	// w.Close()

	// stdin := os.Stdin
	// stdin.Write([]byte("ls\n"))
	// io.WriteString(stdin, "ls\n")
	// _, w := io.Pipe()
	// writer := io.Writer(w)
	// go io.Copy(writer, os.Stdin)
	// go writer.Write([]byte("ls\n"))
	// // stdin.Write(<-stdinChan)
	// stdin.Write([]byte("ls\n"))
	// stdin = stdin.Write(strings.NewReader("ls\n"))
	// io.WriteString(stdin, "")

	// Wait until user exits the shell
	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	// os.Setenv("PROMPT", "()")
	// Keep on keepin' on.
	fmt.Printf("Exited Go Sub Shell\n %s\n", state.String())
}
