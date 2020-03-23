package golang

import (
	"path/filepath"

	"github.com/fatih/color"
	"github.com/magefile/mage/sh"
)

// Lint will call
func Lint(path string) func() error {
	return func() error {
		color.Blue("# Linting source using golangci-lint")
		return sh.RunV("./tools/bin/golangci-lint", "run", "-c", filepath.Clean(filepath.Join(path, ".golangci.yml")), "--fix")
	}
}
