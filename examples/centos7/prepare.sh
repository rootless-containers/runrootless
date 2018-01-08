#!/bin/sh
set -x -e

## 0. clean up
rm -f config.json
[ -d rootfs ] && chmod -R u+w rootfs
rm -rf rootfs

## 1. Generate config,json
runc spec
perl -pi -e 's/"readonly": true/"readonly": false/g' config.json

## 2. Create rootfs
mkdir rootfs
docker pull centos:7
cid=$(docker create centos:7)
# ignore /dev mknod error
docker export $cid | tar Cx ./rootfs 2> /dev/null || true
docker rm -f $cid

## 3. Set up DNS config
# TODO: use bind-mount?
cp -f /etc/hosts /etc/resolv.conf ./rootfs/etc
