#!/bin/sh
set -x -e

cd $(dirname $0)

../common/0-clean-up.sh

../common/1-generate-config-json.sh

## 2. Create rootfs
wget -O /tmp/alpine-minirootfs-3.7.0-x86_64.tar.gz http://dl-cdn.alpinelinux.org/alpine/v3.7/releases/x86_64/alpine-minirootfs-3.7.0-x86_64.tar.gz
# ignore mknod error
tar xzf /tmp/alpine-minirootfs-3.7.0-x86_64.tar.gz -C ./rootfs || true

../common/3-set-up-dns.sh
