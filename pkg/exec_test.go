package pkg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanExecCommand(t *testing.T) {

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := ExecCommand("echo", []string{"arg-1", "arg-2"}, &stdout, &stderr)
	assert.Equal(t, 0, exitCode, "invalid proccess exit code")
	assert.Equal(t, "", stderr.String(), "stderr should be empty")
	assert.Equal(t, "arg-1 arg-2\n", stdout.String(), "stderr should be empty")

}
