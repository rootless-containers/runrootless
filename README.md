# runROOTLESS: rootless OCI container runtime with ptrace hacks

[![Build Status](https://travis-ci.org/AkihiroSuda/runrootless.svg)](https://travis-ci.org/AkihiroSuda/runrootless)

## Quick start (No root privilege is required!)

### Install

Requires: Go, runc

```bash
$ go get github.com/AkihiroSuda/runrootless
$ $GOPATH/src/github.com/AkihiroSuda/runrootless/proot/install.sh
```

Future version should install a pre-built PRoot binary automatically on the first run.

### Usage

Create an example CentOS 7 bundle:

```
$ cd ./examples/centos7
$ ./prepare.sh
$ ls -1F
config.json
prepare.sh
rootfs/
```

Make sure the bundle cannot be executed with the regular `runc`:

```
$ runc run foo
rootless containers require user namespaces
```

Make sure the bundle can be executed with `runrootless`, and you can install some software using `yum`:

```
$ runrootless run foo
sh-4.2# yum install -y epel-release
sh-4.2# yum install -y cowsay
sh-4.2# cowsay hello rootless world
 ______________________
< hello rootless world >
 ----------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## How it works

- Transform a regular `config.json` to rootless one, and create a new OCI runtime bundle with it.
- Bind-mount a static [PRoot](proot) binary so as to allow `apt`/`yum` commands.
- Inject the PRoot binary to `process.args`.
- Invoke plain runC.

## Known issues

- apt-get seems flaky, while yum seems almost stable (CentOS 7). Setting env var `PROOT_NO_SECCOMP=1` may mitigate the issue. (Disables seccomp acceleration)
- apk may not work

## Future work

### OCI Runtime Hook mode

runROOTLESS could be reimplemented as a OCI Runtime Hook (prestart) that works with an arbitrary OCI Runtime.
This work would need adding support for `PTRACE_ATTACH` to PRoot.
Also, it would require YAMA to be disabled.

### Reimplement PRoot in Go

This is hard than I initially thought...

## Legal notice

- [`./proot/PRoot`](./proot/PRoot) originates from [PRoot](https://github.com/proot-me/PRoot) and hence licensed under [GPL v2](./proot/PRoot/COPYING)
- [`./runccompat.go`](./runccompat.go) originates from [runc](https://github.com/opencontainers/runc) (Apache License 2.0)
- Other files are licensed under Apache License 2.0
