package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version injected from CI/CD
	version = `{{.Version}}`

	// Commit hash injected from CI/CD
	commit = `{{.Commit}}`
)

// newVersionCmd return helm-teller command
func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version of the helm teller plugin",
		Run: func(*cobra.Command, []string) {
			fmt.Printf("%s (%s)", version, commit)
		},
	}
}
