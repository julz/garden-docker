# Garden-docker. Two Revolutionary New Products: A docker backend for garden, and a garden front-end for docker.

This is a prototype/proof of concept to see the gaps between garden and docker, and to demonstrate how easy it is to write backends for garden.

# Next Steps

 - Currently we spawn a daemon and ask that to spawn child processes. This is the fastest path from the existing garden-linux architecture to running using docker as a backend. Next we'd like to directly use docker's `exec` command to spawn the processes.
 - Only docker rootfses are supported right now, warden (directory) rootfses won't work. This shouldn't be super hard.
 - Runc runc runc!
 - Disk quotas using btrfs
 - Snapshot/restore
 - ..

# Usage

I wouldn't yet
