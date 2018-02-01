# runROOTLESS: rootless OCI container runtime with ptrace hacks

[![Build Status](https://travis-ci.org/AkihiroSuda/runrootless.svg)](https://travis-ci.org/AkihiroSuda/runrootless)

## Quick start (No root privileges nor SUID binaries are required!)

### Install

Requires: Go, runc

```console
user$ go get github.com/AkihiroSuda/runrootless
user$ $GOPATH/src/github.com/AkihiroSuda/runrootless/install-proot.sh
```

Future version should install a pre-built PRoot binary automatically on the first run.

### Usage

Create an example Ubuntu bundle:

```console
user$ cd ./examples/ubuntu
user$ ./prepare.sh
user$ ls -1F
config.json
prepare.sh
rootfs/
```

Make sure the bundle cannot be executed with the regular `runc`:

```console
user$ runc run ubuntu
rootless containers require user namespaces
```

Note that even with `runc spec --rootless`, you cannot execute `apt`:
```console
user$ rm config.json
user$ runc spec --rootless
user$ sed -i 's/"readonly": true/"readonly": false/' config.json
user$ runc run ubuntu
# apt update
E: setgroups 65534 failed - setgroups (1: Operation not permitted)
E: setegid 65534 failed - setegid (22: Invalid argument)
E: seteuid 100 failed - seteuid (22: Invalid argument)
E: setgroups 0 failed - setgroups (1: Operation not permitted)
Reading package lists... Done
W: chown to _apt:root of directory /var/lib/apt/lists/partial failed - SetupAPTPartialDirectory (22: Invalid argument)
E: setgroups 65534 failed - setgroups (1: Operation not permitted)
E: setegid 65534 failed - setegid (22: Invalid argument)
E: seteuid 100 failed - seteuid (22: Invalid argument)
E: setgroups 0 failed - setgroups (1: Operation not permitted)
E: Method gave invalid 400 URI Failure message: Failed to setgroups - setgroups (1: Operation not permitted)
E: Method http has died unexpectedly!
E: Sub-process http returned an error code (112)_
```

With `runrootless`, you can execute `apt` successfully:

```console
user$ ./prepare.sh
user$ runrootless run ubuntu
# apt update
# apt install -y cowsay
# /usr/games/cowsay hello rootless world
 ______________________
< hello rootless world >
 ----------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

### Other examples

CentOS:
```console
user$ cd ./examples/centos
user$ ./prepare.sh
user$ runrootless run centos
sh-4.2# yum install -y epel-release
sh-4.2# yum install -y cowsay
sh-4.2# cowsay hello rootless world
```

Alpine Linux:
```console
user$ cd ./examples/alpine
user$ ./prepare.sh
user$ runrootless run alpine
/ # apk update
/ # apk add fortune
/ # fortune
```

Arbitrary Docker image:
```console
user$ cd ./examples/docker-image
user$ ./prepare.sh opensuse
user$ runrootless run opensuse
sh-4.3# zypper install cowsay
sh-4.3# cowsay hello rootless world
```

runROOTLESS can be also executed inside Docker container, but `--privileged` is still required ( https://github.com/opencontainers/runc/issues/1456 )

```console
host$ docker run -it --rm --privileged akihirosuda/runrootless
~ $ id
uid=1000(user) gid=1000(user)
~ $ cd ~/examples/ubuntu/
~/examples/ubuntu $ ./prepare.sh
~/examples/ubuntu $ runrootless run ubuntu
#
```

### Environment variables

- `RUNROOTLESS_SECCOMP=1`: enable seccomp acceleration (unstable)

## How it works

- Transform a regular `config.json` to rootless one, and create a new OCI runtime bundle with it.
- Bind-mount a static [PRoot](proot) binary so as to allow `apt`/`yum` commands.
- Inject the PRoot binary to `process.args`.
- Invoke plain runC.

## Known issues

- `apt` / `dpkg` may crash when seccomp acceleration is enabled: https://github.com/AkihiroSuda/runrootless/issues/4

## Future work

### OCI Runtime Hook mode

runROOTLESS could be reimplemented as a OCI Runtime Hook (prestart) that works with an arbitrary OCI Runtime.
This work would need adding support for `PTRACE_ATTACH` to PRoot.
Also, it would require YAMA to be disabled.

### Reimplement PRoot in Go

This is hard than I initially thought...

## Legal notice

- [`./proot/PRoot`](./proot/PRoot) originates from [PRoot](https://github.com/proot-me/PRoot) and hence licensed under [GPL v2](./proot/PRoot/COPYING)
- [`./proot/PRoot/src/rootlesscontainers`](./proot/PRoot/src/rootlesscontainers) (protobuf definition) originates from https://rootlesscontaine.rs/ (Apache License 2.0)
- [`./runccompat.go`](./runccompat.go) originates from [runc](https://github.com/opencontainers/runc) (Apache License 2.0)
- Other files are licensed under Apache License 2.0
