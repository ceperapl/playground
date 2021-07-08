package bundle

import (
	_ "embed"
)

//go:embed bundle.js
var bundle string

func GetBundle() string {
	return bundle
}
