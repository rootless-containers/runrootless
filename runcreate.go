package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rootless-containers/runrootless/bundle"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:   "run",
	Usage:  "create and run a container",
	Flags:  runcRunFlags(),
	Action: runCreate,
}

var createCommand = cli.Command{
	Name:   "create",
	Usage:  "create a container",
	Flags:  runcCreateFlags(),
	Action: runCreate,
}

func runCreate(context *cli.Context) error {
	var err error
	bundleDir := context.String("bundle")
	if bundleDir == "" {
		bundleDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}
	newBundleDir := filepath.Join(bundleDir, "runrootless")
	logrus.Debugf("bundle: %s -> %s", bundleDir, newBundleDir)
	if err := bundle.Transform(newBundleDir, bundleDir); err != nil {
		return err
	}
	redirectToRunc(transformRunCreate(newBundleDir)...)
	return nil
}

// transformRunCreate transforms os.Args for the new bundle.
// e.g. "runc --root /foo run --bundle /bar baz" -> {"--root", "/foo", "run", "baz", "--bundle", newBundle}
func transformRunCreate(newBundle string) []string {
	return _transformRunCreate(os.Args, newBundle)
}

func _transformRunCreate(osArgs []string, newBundle string) []string {
	var ss []string
	runCreate := false
	skipNext := false
	for _, s := range osArgs[1:] {
		if skipNext {
			skipNext = false
			continue
		}
		if s == "run" || s == "create" {
			runCreate = true
		}
		if runCreate {
			if s == "-b" || s == "--bundle" {
				skipNext = true
				continue
			} else if strings.HasPrefix(s, "-b=") || strings.HasPrefix(s, "--bundle=") {
				continue
			}
		}
		ss = append(ss, s)
	}
	if runCreate {
		ss = append(ss, "--bundle", newBundle)
	}
	return ss
}
