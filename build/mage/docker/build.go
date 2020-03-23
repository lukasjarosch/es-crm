package docker

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	"github.com/lukasjarosch/escrm/build/mage/git"
)

// Command specification for dockerfile generation.
type Command struct {
	Service     string
	Name        string
	Description string
	URL         string
}

// Generate Dockerfile for given command.
func Generate(cmd *Command) func() error {
	return func() error {
		buf, err := merge(cmd)
		if err != nil {
			return err
		}

		// Write output to Stdout
		_, errWrite := buf.WriteTo(os.Stdout)
		return errWrite
	}
}

// Build a docker container for given command.
func Build(cmd *Command) func() error {
	return func() error {
		mg.Deps(git.CollectInfo)

		buf, err := merge(cmd)
		if err != nil {
			return err
		}

		// Invoke docker commands
		err = sh.RunWith(
			map[string]string{
				"DOCKER_BUILDKIT": "1",
			},
			"/bin/sh", "-c",
			fmt.Sprintf(
				"echo '%s' | base64 -d | docker build -t sax-platform/%s:v%s -f- --build-arg BUILD_DATE=%s --build-arg VERSION=%s --build-arg VCS_REF=%s .",
				base64.StdEncoding.EncodeToString(buf.Bytes()),
				cmd.Service,
				git.Revision,
				time.Now().Format(time.RFC3339),
				git.Tag,
				git.Revision,
			),
		)

		return err
	}
}

// -----------------------------------------------------------------------------

func merge(cmd *Command) (*bytes.Buffer, error) {

	data, err := ioutil.ReadFile("./build/mage/docker/Dockerfile")
	if err != nil {
		return nil, err
	}

	// Compile template
	dockerFileTmpl, err := template.New("Dockerfile").Parse(string(data))
	if err != nil {
		return nil, err
	}

	// Merge data
	var buf bytes.Buffer
	if errTmpl := dockerFileTmpl.Execute(&buf, cmd); errTmpl != nil {
		return nil, errTmpl
	}

	// Return buffer without error
	return &buf, nil
}
