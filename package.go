//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target . -package policeuk --clean openapi.yaml
package policeuk

import (
	_ "github.com/ogen-go/ogen/gen"
)

const (
	ServerURL = "https://data.police.uk/api"
)
