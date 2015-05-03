// initd runs as pid 1 inside a container, listens on a socket and spawns processes
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudfoundry-incubator/garden-linux/container_daemon"
	"github.com/cloudfoundry-incubator/garden-linux/container_daemon/unix_socket"
	"github.com/cloudfoundry-incubator/garden-linux/containerizer/system"
	"github.com/pivotal-golang/lager"
)

func main() {
	logger := lager.NewLogger("initd")
	socketPath := flag.String("socketPath", "/run/initd.sock", "path to listen for spawn requests on")
	//unmountPath := flag.String("unmountAfterListening", "/run", "directory to unmount after succesfully listening on -socketPath")
	flag.String("unmountAfterListening", "/run", "directory to unmount after succesfully listening on -socketPath")
	flag.Parse()

	reaper := system.StartReaper(logger)
	defer reaper.Stop()

	listener := &unix_socket.Listener{SocketPath: *socketPath}

	daemon := container_daemon.ContainerDaemon{
		Listener: listener,
		Users:    &system.LibContainerUser{},
		Runner:   reaper,
	}

	// open up the listener socket
	if err := listener.Init(); err != nil {
		fmt.Printf("listen on %s: %s", *socketPath, err)
		os.Exit(1)
	}

	// unmount the bind-mounted socket volume now we've started listening
	// if err := syscall.Unmount(*unmountPath, syscall.MNT_DETACH); err != nil {
	// 	fmt.Printf("unmount %s: %s", *unmountPath, err)
	// 	os.Exit(2)
	// }

	// let daemon take over and start listening for incoming start process requests
	if err := daemon.Run(); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
