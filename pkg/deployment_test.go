package pkg

import (
	"helm-teller/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseToSetFlags(t *testing.T) {

	entries := map[string]string{
		"key-1": "val-1",
		"key-2": "val-2",
	}
	teller := testutils.GetTellerPkg(entries)
	args, returnEntries, err := ParseToSetFlags(teller, "install")
	assert.Nil(t, err)
	assert.ElementsMatch(t, []string{"install", "--set", "teller.key-1=\"val-1\"", "--set", "teller.key-2=\"val-2\""}, args)
	assert.Equal(t, entries, returnEntries)

}

func TestParseToSetFlagsWithCollectErr(t *testing.T) {

	entries := map[string]string{}
	teller := testutils.GetTellerPkg(entries)
	args, returnEntries, err := ParseToSetFlags(teller, "install")
	assert.NotNil(t, err)
	assert.Nil(t, args)
	assert.Nil(t, returnEntries)

}

func TestParseToSetFlagsWithExportYAMLErr(t *testing.T) {

	entries := map[string]string{"error": "error"}
	teller := testutils.GetTellerPkg(entries)
	args, returnEntries, err := ParseToSetFlags(teller, "install")
	assert.NotNil(t, err)
	assert.Nil(t, args)
	assert.Nil(t, returnEntries)

}

func TestParseToSetFlagsWithParsingYamlErr(t *testing.T) {

	entries := map[string]string{
		"invalid-yaml": "val-1",
	}
	teller := testutils.GetTellerPkg(entries)
	args, returnEntries, err := ParseToSetFlags(teller, "install")
	assert.NotNil(t, err)
	assert.Nil(t, args)
	assert.Nil(t, returnEntries)

}

func TestMaskHelmOutput(t *testing.T) {
	str := `
apiVersion: v1
kind: ConfigMap
metadata:
	name: test-config-map
data:
	redis-host: localhost
	redis-password: 1234
	loglevel: debug	
`
	expectedStr := "\napiVersion: v1\nkind: ConfigMap\nmetadata:\n\tname: test-config-map\ndata:\n\tredis-host: lo*****\n\tredis-password: 12*****\n\tloglevel: debug\t\n"
	entries := map[string]string{
		"key-1": "localhost",
		"key-2": "1234",
	}
	out := MaskHelmOutput(str, entries)
	assert.Equal(t, expectedStr, out)
}

func TestParseSetEntry(t *testing.T) {

	args := parseSetEntry("key", "value")
	assert.ElementsMatch(t, []string{"--set", "teller.key=\"value\""}, args)
}
