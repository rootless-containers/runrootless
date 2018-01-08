package bundle

import (
	"os"
	"path/filepath"

	"github.com/opencontainers/runc/libcontainer/specconv"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
)

func transformSpec(spec *specs.Spec, oldBundle string) error {
	specconv.ToRootless(spec)
	toAbsoluteRootFS(spec, oldBundle)
	return injectPRoot(spec)
}

func toAbsoluteRootFS(spec *specs.Spec, oldBundle string) {
	if !filepath.IsAbs(spec.Root.Path) {
		spec.Root.Path = filepath.Clean(filepath.Join(oldBundle, spec.Root.Path))
	}
}

func injectPRoot(spec *specs.Spec) error {
	proot, err := prootPath()
	if err != nil {
		return err
	}
	spec.Mounts = append(spec.Mounts,
		specs.Mount{
			Destination: "/dev/proot",
			Type:        "tmpfs",
			Source:      "tmpfs",
			Options:     []string{"exec", "mode=755", "size=32256k"},
		},
		specs.Mount{
			Destination: "/dev/proot/proot",
			Type:        "bind",
			Source:      proot,
			Options:     []string{"bind", "ro"},
		},
	)
	spec.Process.Args = append([]string{"/dev/proot/proot", "-0"}, spec.Process.Args...)
	spec.Process.Env = append(spec.Process.Env, "PROOT_TMP_DIR=/dev/proot")
	return nil
}

func prootPath() (string, error) {
	// we can't use os/user.Current in a static binary.
	// moby/moby#29478
	home := os.Getenv("HOME")
	s := filepath.Join(home, ".runrootless", "runrootless-proot")
	_, err := os.Stat(s)
	if os.IsNotExist(err) {
		return s, errors.Errorf("%s not found. please install runrootless-proot according to README.", s)
	}
	return s, err
}
