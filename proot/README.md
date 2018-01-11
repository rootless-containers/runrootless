# PRoot for runROOTLESS

The [`PRoot`](PRoot) directory contains a fork of PRoot.
Please refer to [`PRoot/COPYING`](PRoot/COPYING) for the license information (GPL v2).

Please use `install.sh` to install.

Note that the upstream version of PRoot ( https://github.com/proot-me/PRoot) is unlikely to work with runROOTLESS.

## Changes from the PRoot upstream

### patches from udocker

https://github.com/jorge-lip/PRoot/commit/10ca3e88dc1d2e2b45439b181a168af6b4053b91

### persistent chown

Following xattrs are used for storing decimal uid/gid strings:

- `user.rootlesscontainers.uid`
- `user.rootlesscontainers.gid`

### support for chroot(relative_path)

Currently, only chroot to the current root is supported.
(This hack is required for running `apk`)
