#!/bin/sh
set -x -e
wget -O rootfs.tar.gz http://dl-cdn.alpinelinux.org/alpine/v3.7/releases/x86_64/alpine-minirootfs-3.7.0-x86_64.tar.gz
rm -rf rootfs
mkdir rootfs
# ignore mknod error
tar xzf rootfs.tar.gz -C ./rootfs || true
rm -f rootfs.tar.gz
