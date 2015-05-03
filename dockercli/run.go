package dockercli

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cloudfoundry/gunk/command_runner"
)

type Runner struct {
	Runner command_runner.CommandRunner
}

func (r *Runner) Run(cmd RunCmd) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c := cmd.Cmd()
	c.Stdout = &stdout
	c.Stderr = &stderr

	if err := r.Runner.Run(c); err != nil {
		return "", fmt.Errorf("run: %s: %s", err, strings.TrimRight(stderr.String(), "\n"))
	}

	return strings.TrimRight(stdout.String(), "\n"), nil
}

func (r *Runner) Inspect(cmd InspectCmd) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	c := cmd.Cmd()
	c.Stdout = &stdout
	c.Stderr = &stderr

	if err := r.Runner.Run(c); err != nil {
		return "", fmt.Errorf("inspect: %s: %s", err, strings.TrimRight(stderr.String(), "\n"))
	}

	return strings.TrimRight(stdout.String(), "\n"), nil
}
