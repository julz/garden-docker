package dockercli

import (
	"fmt"
	"os/exec"
)

type RunCmd struct {
	Volumes []Volume
	Image   string

	Program     string
	ProgramArgs []string
	Detach      bool
}

type Volume struct {
	HostPath      string
	ContainerPath string
}

func (cmd *RunCmd) Cmd() *exec.Cmd {
	program := append([]string{cmd.Program}, cmd.ProgramArgs...)
	volumes := []string{}
	for _, v := range cmd.Volumes {
		volumes = append(volumes, "-v", v.arg())
	}

	args := append(append(volumes, cmd.Image), program...)

	if cmd.Detach {
		args = append([]string{"-d"}, args...)
	}

	return exec.Command("docker", append([]string{"run"}, args...)...)
}

func (v Volume) arg() string {
	return fmt.Sprintf("%s:%s", v.HostPath, v.ContainerPath)
}

type InspectCmd struct {
	ContainerID string
	Field       string
}

func (cmd *InspectCmd) Cmd() *exec.Cmd {
	return exec.Command("docker", "inspect", fmt.Sprintf("--format='{{.%s}}'", cmd.Field), cmd.ContainerID)
}
