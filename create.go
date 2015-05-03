package gardendocker

import (
	"fmt"
	"net/url"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/garden-linux/old/port_pool"
	"github.com/cloudfoundry-incubator/garden-linux/process_tracker"
	"github.com/cloudfoundry/gunk/command_runner"
	"github.com/docker/docker/pkg/iptables"
	"github.com/julz/garden-docker/dockercli"
)

type DaemonContainerCreator struct {
	DefaultRootfs string
	Depot         Depot

	DoshPath  string
	InitdPath string

	Chain    *iptables.Chain
	PortPool *port_pool.PortPool

	DockerRunner  DockerRunner
	CommandRunner command_runner.CommandRunner
}

//go:generate counterfeiter . DockerRunner
type DockerRunner interface {
	Run(dockercli.RunCmd) (string, error)
	Inspect(dockercli.InspectCmd) (string, error)
}

func (c *DaemonContainerCreator) Create(spec garden.ContainerSpec) (*Container, error) {
	dir, err := c.Depot.Create()
	if err != nil {
		return nil, fmt.Errorf("create depot dir: %s", err)
	}

	if len(spec.RootFSPath) == 0 {
		spec.RootFSPath = c.DefaultRootfs
	}

	rootfs, err := url.Parse(spec.RootFSPath)
	if err != nil {
		return nil, fmt.Errorf("create: not a valid rootfs path: %s", err)
	}

	var dockerID string
	if dockerID, err = c.DockerRunner.Run(dockercli.RunCmd{
		Image:       rootfs.Path[1:],
		Detach:      true,
		Program:     "/garden-bin/initd",
		ProgramArgs: []string{"-socketPath", "/run/initd.sock", "-unmountAfterListening", "/run"},
		Volumes: []dockercli.Volume{
			{
				HostPath:      c.InitdPath,
				ContainerPath: "/garden-bin/initd",
			},
			{
				HostPath:      path.Join(dir, "run"),
				ContainerPath: "/run",
			},
		},
	}); err != nil {
		return nil, fmt.Errorf("create: %s", err)
	}

	var ip string
	if ip, err = c.DockerRunner.Inspect(dockercli.InspectCmd{
		ContainerID: dockerID,
		Field:       "NetworkSettings.IPAddress",
	}); err != nil {
		return nil, fmt.Errorf("create: inspect %s: %s", dockerID, err)
	}

	return &Container{
		LimitsHandler: &LimitsHandler{},
		StreamHandler: &StreamHandler{},
		InfoHandler: &InfoHandler{
			Spec:          spec,
			ContainerPath: dir,
			ContainerIP:   ip,
			DockerID:      dockerID,
			PropsHandler:  &PropsHandler{},
		},
		NetHandler: &NetHandler{
			ContainerIP: ip,
			Chain:       c.Chain,
			PortPool:    c.PortPool,
		},
		RunHandler: &RunHandler{
			ProcessTracker: process_tracker.New(dir, c.CommandRunner),
			ContainerCmd: &doshcmd{
				Path:      filepath.Join(dir, "bin", "dosh"),
				InitdSock: filepath.Join(dir, "run", "initd.sock"),
			},
		},
	}, nil
}

type doshcmd struct {
	Path      string
	InitdSock string
}

func (d doshcmd) Cmd(path string, args ...string) *exec.Cmd {
	doshArgs := []string{"-socketPath", d.InitdSock, "-user", "root"}
	run := []string{path}
	run = append(run, args...)
	return exec.Command(d.Path, append(doshArgs, run...)...)
}
