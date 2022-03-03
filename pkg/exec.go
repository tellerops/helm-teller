package pkg

import (
	"io"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func ExecCommand(name string, args []string, stdout, stderr io.Writer) int {

	helmCmd := exec.Command(name, args...)
	helmCmd.Stdout = stdout
	helmCmd.Stderr = stderr
	if err := helmCmd.Run(); err != nil {
		log.WithError(err).Trace("could not execute command")
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}
	return 0
}
