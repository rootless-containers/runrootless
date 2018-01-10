#!/bin/sh
set -x -e

cd $(dirname $0)

../common/0-clean-up.sh

../common/1-generate-config-json.sh

## 2. Create rootfs
wget -O /tmp/centos-7-x86_64-minimal.tar.gz https://download.openvz.org/template/precreated/centos-7-x86_64-minimal.tar.gz
# ignore mknod error
tar xzf /tmp/centos-7-x86_64-minimal.tar.gz -C ./rootfs || true

../common/3-set-up-dns.sh
