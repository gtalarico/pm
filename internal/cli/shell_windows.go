package cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"
	"golang.org/x/sys/windows"
)

// unsafe bindings to system dlls to find parent process path
// see getParentProcessPath
var (
	modkernel32 = windows.NewLazySystemDLL("Psapi.dll")
	syscallName = modkernel32.NewProc("GetModuleFileNameExA")
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

// Find path of parent process
// https://stackoverflow.com/a/1933140/2350244
func getParentProcessPath() (path string, err error) {
	// open parent process
	handle, err := windows.OpenProcess(
		windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ,
		false,
		uint32(os.Getppid()),
	)
	if err != nil {
		return "", errors.New("Can not access parent process")
	}

	// make a syscall to get path of parent process into buffer
	// Note: buffer is ascii. this method returns an LCPTR in windows
	filename := make([]uint8, windows.MAX_LONG_PATH)
	// similar to source in
	// https://pkg.go.dev/golang.org/x/sys@v0.0.0-20200930185726-fdedc70b468f/windows#GetModuleHandleEx
	// syscall6 takes 6 parameters, wrapping a native system call
	// we are passing 4 arguments so the last 2 are null
	// close parent process handle when done
	defer windows.CloseHandle(handle)
	r0, _, e1 := syscall.Syscall6(
		syscallName.Addr(),                    // address of api method
		4,                                     // 4 inputs
		uintptr(handle),                       // process handle
		0,                                     // module handle, leave empty to get process exe path
		uintptr(unsafe.Pointer(&filename[0])), // buffer pointer
		uintptr(windows.MAX_PATH),             // max path size
		0, 0,                                  // null values
	)
	// process results
	res := uint32(r0)
	if e1 != 0 || res == 0 {
		return "", errors.New("Can not access parent process")
	}

	// remove all 0s from buffer
	b := make([]byte, 0)
	for _, v := range filename {
		if byte(v) != 0 {
			b = append(b, byte(v))
		} else {
			break
		}
	}

	// convert buffer into string
	return string(b), nil
}

func Shell(cwd string) {
	//technosophos.com/2014/07/11/start-an-interactive-shell-from-within-go.html
	// defer handleShellError()

	// silent the ctrl+c / SIGTERM signals
	silentCtrlC()

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	// Start up a new shell
	shellPath, err := getParentProcessPath()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Starting new shell (%s)\n", shellPath)
	fmt.Println("Use 'exit' to terminate child shell")
	proc, err := os.StartProcess(shellPath, []string{}, &pa)
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
