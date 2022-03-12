package cmd

import (
	"bytes"
	"fmt"
	"helm-teller/pkg"
	"os"

	"github.com/spf13/cobra"
)

type deployCmd struct {
	show           bool
	disableMasking bool
}

// newDeploymentCommand creates a helm wrapper deployment command
func newDeploymentCommand(deployType string) *cobra.Command {

	deployCmd := deployCmd{}

	cmd := &cobra.Command{
		Use:     deployType,
		Short:   fmt.Sprintf("wrapper for %s helm command", deployType),
		PreRunE: loadTellerFile,
		Run: func(cmd *cobra.Command, args []string) {

			helmCustomFlags, entries, err := pkg.ParseToSetFlags(teller, deployType)
			if err != nil {
				os.Exit(1)
			}

			if deployCmd.show {
				teller.PrintEnvKeys()
			}

			helmCustomFlags = append(helmCustomFlags, args...)

			var stdout bytes.Buffer
			var stderr bytes.Buffer
			exitCode := pkg.ExecCommand(helmBinary, helmCustomFlags, &stdout, &stderr)

			if exitCode != 0 {
				fmt.Println(stderr.String())
			}
			if deployCmd.disableMasking {
				fmt.Println(stdout.String())
			} else {
				fmt.Println(pkg.MaskHelmOutput(stdout.String(), entries))
			}

			os.Exit(exitCode)
		},
	}

	f := cmd.Flags()
	f.BoolVar(&deployCmd.show, "show", false, "Print in a human friendly, secure format.")
	f.BoolVar(&deployCmd.disableMasking, "disable-masking", false, "Disable configuration masking.")

	return cmd

}
