// dosh connects to a running garden-docker container daemon, spawn a process, stream output
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/garden-linux/container_daemon"
	"github.com/cloudfoundry-incubator/garden-linux/container_daemon/unix_socket"

	_ "github.com/cloudfoundry-incubator/garden-linux/iodaemon"
)

func main() {
	socketPath := flag.String("socketPath", "./run/initd.sock", "socket initd is listening on")
	dir := flag.String("dir", "", "working directory for spawned process")
	user := flag.String("user", "", "user to run container as (defaults to current user)")

	flag.Parse()

	extraArgs := flag.Args()
	if len(extraArgs) == 0 {
		fmt.Fprintf(os.Stderr, "Command name not provided.")
		os.Exit(container_daemon.UnknownExitStatus)
	}

	processSpec := &garden.ProcessSpec{
		Path: extraArgs[0],
		Args: extraArgs[1:],
		Env:  []string{"HELLO=1"},
		Dir:  *dir,
		User: *user,
	}

	processIO := &garden.ProcessIO{
		Stdin:  os.Stdin,
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}

	connector := &unix_socket.Connector{
		SocketPath: *socketPath,
	}

	proc, err := container_daemon.NewProcess(connector, processSpec, processIO)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Starting process: %s", err)
		os.Exit(container_daemon.UnknownExitStatus)
	}

	exitCode, err := proc.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Waiting for process to complete: %s", err)
		os.Exit(container_daemon.UnknownExitStatus)
	}

	os.Exit(exitCode)
}
