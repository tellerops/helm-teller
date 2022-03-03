package pkg

import (
	"fmt"
	"math"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type TellerPkgDescriber interface {
	Collect() error
	ExportYAML() (out string, err error)
	PrintEnvKeys()
}

// ParseToSetFlags will collect teller configuration and return a list of --set flag for each path
func ParseToSetFlags(teller TellerPkgDescriber, helmCommand string) ([]string, map[string]string, error) {

	err := teller.Collect()
	if err != nil {
		log.WithError(err).Error("could not collect teller secrets")
		return nil, nil, err
	}

	out, err := teller.ExportYAML()
	if err != nil {
		log.WithError(err).Error("error while export teller config to yaml")
		return nil, nil, err
	}

	entries := map[string]string{}
	err = yaml.Unmarshal([]byte(out), &entries)
	if err != nil {
		log.WithError(err).Error("could not parse teller entries")
		return nil, nil, err
	}

	tellerSets := []string{}

	for key, val := range entries {
		tellerSets = append(tellerSets, parseSetEntry(key, val)...)
	}

	command := []string{helmCommand}
	command = append(command, tellerSets...)

	return command, entries, nil
}

// MaskHelmOutput replace given entries string to mask chars in the given str
func MaskHelmOutput(str string, entries map[string]string) string {

	for _, value := range entries {
		str = strings.Replace(str, value, fmt.Sprintf("%s*****", value[:int(math.Min(float64(len(value)), 2))]), -1)
	}
	return str
}

// parseSetEntry return --set flag for the given key value
func parseSetEntry(key, value string) []string {

	return []string{"--set", fmt.Sprintf(`teller.%s="%s"`, key, value)}
}
