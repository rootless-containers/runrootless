package main

import (
	"os"

	"github.com/urfave/cli"
)

// runcGlobalFlags original: runc@e6516b3d5dc780cb57a976013c242a9a93052543, main.go
func runcGlobalFlags() []cli.Flag {
	root := "/run/runc"
	if os.Geteuid() != 0 {
		runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
		if runtimeDir != "" {
			root = runtimeDir + "/runc"
		}
	}
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output for logging",
		},
		cli.StringFlag{
			Name:  "log",
			Value: "/dev/null",
			Usage: "set the log file path where internal debug information is written",
		},
		cli.StringFlag{
			Name:  "log-format",
			Value: "text",
			Usage: "set the format used by logs ('text' (default), or 'json')",
		},
		cli.StringFlag{
			Name:  "root",
			Value: root,
			Usage: "root directory for storage of container state (this should be located in tmpfs)",
		},
		cli.StringFlag{
			Name:  "criu",
			Value: "criu",
			Usage: "path to the criu binary used for checkpoint and restore",
		},
		cli.BoolFlag{
			Name:  "systemd-cgroup",
			Usage: "enable systemd cgroup support, expects cgroupsPath to be of form \"slice:prefix:name\" for e.g. \"system.slice:runc:434234\"",
		},
	}
}

// runcCreateFlags original: runc@e6516b3d5dc780cb57a976013c242a9a93052543, create.go
func runcCreateFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "bundle, b",
			Value: "",
			Usage: `path to the root of the bundle directory, defaults to the current directory`,
		},
		cli.StringFlag{
			Name:  "console-socket",
			Value: "",
			Usage: "path to an AF_UNIX socket which will receive a file descriptor referencing the master end of the console's pseudoterminal",
		},
		cli.StringFlag{
			Name:  "pid-file",
			Value: "",
			Usage: "specify the file to write the process id to",
		},
		cli.BoolFlag{
			Name:  "no-subreaper",
			Usage: "disable the use of the subreaper used to reap reparented processes",
		},
		cli.BoolFlag{
			Name:  "no-pivot",
			Usage: "do not use pivot root to jail process inside rootfs.  This should be used whenever the rootfs is on top of a ramdisk",
		},
		cli.BoolFlag{
			Name:  "no-new-keyring",
			Usage: "do not create a new session keyring for the container.  This will cause the container to inherit the calling processes session key",
		},
		cli.IntFlag{
			Name:  "preserve-fds",
			Usage: "Pass N additional file descriptors to the container (stdio + $LISTEN_FDS + N in total)",
		},
	}
}

// runcRunFlags original: runc@e6516b3d5dc780cb57a976013c242a9a93052543, run.go
func runcRunFlags() []cli.Flag {
	return append(runcCreateFlags(),
		cli.BoolFlag{
			Name:  "detach, d",
			Usage: "detach from the container's process",
		})
}
