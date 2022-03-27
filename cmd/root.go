package cmd

import (
	log "github.com/sirupsen/logrus"
	tellerPkg "github.com/spectralops/teller/pkg"
	"github.com/spf13/cobra"

	"github.com/SpectralOps/helm-teller/pkg"
	"github.com/SpectralOps/helm-teller/pkg/visibility"
)

var (
	// logLevel define the application log level
	logLevel string

	// helmBinary define the binary helm path
	helmBinary string

	// tellerConfig point to Teller configuration file
	tellerConfig string

	// teller descrive Teller package
	teller pkg.TellerPkgDescriber
)

const rootCmdLongUsage = `
Helm Teller Allows you to inject configuration and secrets from multiple providers into your chart while masking the secrets at the deployment.

* More secure while using --debug or --dry-run the secrets will not show in the STDOUT.
* Simple to integrate.
* Rich of supported plugins.
* Pull configuration and secret from multiple providers in one place.
* Manage configuration from development to production in the same way.
`

func New() *cobra.Command {

	cobra.OnInitialize(initialize)

	rootCmd := &cobra.Command{
		Use:   "helm-teller",
		Short: "Collect configuration from multiple provider",
		Long:  rootCmdLongUsage,
	}

	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "Application log level")
	rootCmd.PersistentFlags().StringVar(&helmBinary, "helm-binary", "helm", "Helm binary path")
	rootCmd.PersistentFlags().StringVar(&tellerConfig, "teller-config", ".teller.yaml", "Path to teller.yml config")

	rootCmd.AddCommand(newVersionCmd())                 // add version command
	rootCmd.AddCommand(newDeploymentCommand("upgrade")) // add upgrade command
	rootCmd.AddCommand(newDeploymentCommand("install")) // add install command

	return rootCmd
}

// initialize app on each command's
func initialize() {
	// init logger
	visibility.SetLoggingLevel(logLevel)

}

func loadTellerFile(cmd *cobra.Command, args []string) error {
	// init teller package from config file.
	tlrfile, err := tellerPkg.NewTellerFile(tellerConfig)
	if err != nil {
		log.WithError(err).Fatal("could not load teller file")
		return err
	}

	teller = tellerPkg.NewTeller(tlrfile, []string{}, false)
	return nil
}
