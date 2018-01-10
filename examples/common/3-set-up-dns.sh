#!/bin/sh
set -x -e

## 3. Set up DNS config
# TODO: use bind-mount?
cp -f /etc/hosts /etc/resolv.conf ./rootfs/etc
