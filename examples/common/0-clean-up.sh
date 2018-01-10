#!/bin/sh
set -x -e

## 0. clean up
rm -f config.json
[ -d rootfs ] && chmod -R u+w rootfs
rm -rf rootfs
mkdir rootfs
