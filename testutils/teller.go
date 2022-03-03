package testutils

import (
	"bytes"
	"errors"
	"fmt"
)

type MockTellerPkg struct {
	tellerEntries map[string]string
}

func GetTellerPkg(tellerEntries map[string]string) *MockTellerPkg {

	return &MockTellerPkg{
		tellerEntries: tellerEntries,
	}
}

func (mt *MockTellerPkg) Collect() error {

	if len(mt.tellerEntries) == 0 {
		return errors.New("empty entries")
	}
	return nil
}

func (mt *MockTellerPkg) ExportYAML() (out string, err error) {

	var b bytes.Buffer

	for k, v := range mt.tellerEntries {
		if k == "error" {
			return "", errors.New("ExportYAML error")
		}
		if k == "invalid-yaml" {
			b.WriteString(fmt.Sprintf("%s - %s", k, v))
			continue
		}
		b.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	return b.String(), nil
}

func (mt *MockTellerPkg) PrintEnvKeys() {

}
