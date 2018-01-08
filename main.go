package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	runc = "runc"
)

func main() {
	app := cli.NewApp()
	app.Name = "runrootless"
	app.Usage = "rootless OCI container runtime. CLI is runc-compatible."
	app.Flags = runcGlobalFlags()
	app.CommandNotFound = onCommandNotFound
	app.Commands = []cli.Command{
		runCommand,
		createCommand,
	}
	cli.VersionPrinter = printVersion
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func onCommandNotFound(c *cli.Context, s string) {
	logrus.Debugf("unimplemented command: %q. redirecting to runc.", s)
	redirectToRunc()
}

func redirectToRunc(args ...string) {
	if len(args) == 0 {
		args = os.Args[1:]
	}
	cmd := exec.Command(runc, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if ws, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				os.Exit(ws.ExitStatus())
			}
		}
		logrus.Error(err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func printVersion(c *cli.Context) {
	redirectToRunc("--version")
	fmt.Fprintf(c.App.Writer, "%v version %v\n", c.App.Name, c.App.Version)
}
