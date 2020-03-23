package golang

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	"github.com/lukasjarosch/escrm/build/mage/git"
)

func Build(packageName, binaryName string) func() error {
	return func() error {
		mg.Deps(git.CollectInfo)

		version, err := git.TagMatch(fmt.Sprintf("%s*", binaryName))
		if err != nil {
			return err
		}

		fmt.Printf("=> building binary '%s' version %s [%s]\n", binaryName, version, packageName)

		// setup linker args to be used as LDFLAGs
		var args []string
		linkerMap := map[string]string{
			"github.com/lukasjarosch/escrm/build/version.Version":   version,
			"github.com/lukasjarosch/escrm/build/version.Revision":  git.Revision,
			"github.com/lukasjarosch/escrm/build/version.Branch":    git.Branch,
			"github.com/lukasjarosch/escrm/build/version.User":      os.Getenv("USER"),
			"github.com/lukasjarosch/escrm/build/version.Date":      time.Now().Format(time.RFC3339),
			"github.com/lukasjarosch/escrm/build/version.GoVersion": runtime.Version(),
		}
		for name, value := range linkerMap {
			args = append(args, "-X", fmt.Sprintf("%s=%s", name, value))
		}
		args = append(args, "-s", "-w")

		// build artifact
		buildEnv := map[string]string{
			"CGO_ENABLED": "0",
			"GOOS":        "linux",
			"GOARCH":      "amd64",
		}
		binaryPath := fmt.Sprintf("../../bin/%s", binaryName)
		err = sh.RunWith(buildEnv, "go", "build", "-ldflags", strings.Join(args, " "), "-mod=vendor", "-o", binaryPath, packageName)
		if err != nil {
			return err
		}

		path, _ := filepath.Abs(binaryPath)
		fmt.Printf("=> binary compiled into: %s\n", path)
		return nil
	}
}
