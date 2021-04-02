package shell

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/nobe4/safe/internal/logger"
)

const (
	NOTICE_ENTER = `Entering safe space.
  Press Ctrl-c to exit.
Run the following to remember you're in a safe space:
  PS1="(safe) $PS1"`
	NOTICE_EXIT = `Exiting safe space.`
)

// isShell decide to start a new shell or to use the piped data.
// true means starts a new shell, false means use the piped data
// adatped from https://coderwall.com/p/zyxyeg/golang-having-fun-with-os-stdin-and-shell-pipes
func IsShell() (bool, error) {
	logger.Debug("Check if running in a shell.")

	logger.Debug("Check os.Stdin stat")

	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}

	logger.Debug("Use os.Stdin mode.")

	return (fi.Mode()&os.ModeNamedPipe == 0), nil
}

// getShell get the currently running shell.
func getShell() string {
	logger.Debug("Check currently running shell.")

	logger.Debug("Check $SHELL")

	sh := os.Getenv("SHELL")
	if sh != "" {
		logger.Debug("Found $SHELL: " + sh)

		return sh
	}

	// Shells to test for, by order of priority.
	shells := []string{
		"zsh", "bash", "sh",
	}

	for _, s := range shells {
		logger.Debug("Check " + s + " path")

		sh, _ = exec.LookPath(s)
		if sh != "" {
			logger.Debug("Found " + s + ": " + sh)

			return sh
		}
	}

	logger.Warn("Found no shell")

	return ""
}

// Start starts the shell and uses the writer to filter strings.
func Start(writer io.Writer) error {
	fmt.Println(NOTICE_ENTER)

	logger.Debug("Start the shell")

	sh := getShell()
	if sh == "" {
		return errors.New("nil shell")
	}

	logger.Debug("Create command with shell: " + sh)

	cmd := exec.Command(sh)
	if cmd == nil {
		return fmt.Errorf("nil command for shell %s", sh)
	}

	logger.Debug("Assign the pipes.")

	cmd.Stdin = os.Stdin
	cmd.Stdout = writer
	cmd.Stderr = os.Stderr

	logger.Debug("Start the command")

	if err := cmd.Start(); err != nil {
		return err
	}

	logger.Debug("Wait for the command to end.")

	err := cmd.Wait()

	fmt.Println(NOTICE_EXIT)

	return err
}
