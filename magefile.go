// +build mage

package main

import (
	"github.com/lukasjarosch/escrm/build/mage/golang"
	"github.com/lukasjarosch/escrm/build/mage/tools"
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type API mg.Namespace

// Generates the API stub implementations from the protobuf definitions using 'protoc'
func (API) Generate() error {
	mg.Deps(tools.Env)

	protoPath := "./proto/crm"

	color.White("=> generating protobuf stubs in '%s' using prototool", protoPath)
	if err := sh.RunV("./tools/bin/prototool", "all", "--fix", protoPath); err != nil {
		color.Red("=> %s", err)
		return err
	}
	return nil
}

type Code mg.Namespace

func (Code) Lint() error {
	mg.Deps(golang.Lint("."))
	return nil
}
