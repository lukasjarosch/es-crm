package golang

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Vendor() error {
	fmt.Println("=> vendoring dependencies")
	return sh.RunV("go", "mod", "vendor")
}
