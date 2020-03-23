// +build mage

package main

import (
	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

func Build() {
	color.Blue("# Installing third-party tools")
	mg.Deps(Go.Vendor)
	mg.Deps(Go.Tools)
	color.Blue("# Binaries compiled into ./bin")
}

type Go mg.Namespace

func (Go) Vendor() error {
	color.Yellow("+ vendoring tool dependencies")
	return sh.RunV("go", "mod", "vendor")
}

func (Go) Tools() error {
	color.Yellow("+ installing tools using gex")
	return sh.RunV("go", "run", "github.com/izumin5210/gex/cmd/gex", "--build")
}


