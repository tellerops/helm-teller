package main

import (
	"os"

	"github.com/SpectralOps/helm-teller/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		log.WithError(err)
		os.Exit(1)
	}
	os.Exit(0)
}
