package templates

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../transactions/... -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../transactions/...

import (
	"strings"

	"github.com/onflow/flow-core-contracts/lib/go/templates/internal/assets"
)

const (
	filePath             = "../../../transactions/"
	inspectFieldFilename = "inspect_field.cdc"
	defaultField         = "transactionField"
)

// GenerateInspectFieldScript creates a script that reads
// a field from the smart contract and makes assertions
// about its value
func GenerateInspectFieldScript(field string) []byte {
	code := assets.MustAssetString(filePath + inspectFieldFilename)

	code = strings.ReplaceAll(
		code,
		defaultField,
		field,
	)

	return []byte(code)
}
