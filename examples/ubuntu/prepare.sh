#!/bin/sh
set -x -e

cd $(dirname $0)

../common/0-clean-up.sh

../common/1-generate-config-json.sh

## 2. Create rootfs
wget -O /tmp/ubuntu-base-17.10-base-amd64.tar.gz http://cdimage.ubuntu.com/ubuntu-base/releases/17.10/release/ubuntu-base-17.10-base-amd64.tar.gz
# ignore mknod error
tar xzf /tmp/ubuntu-base-17.10-base-amd64.tar.gz -C ./rootfs || true

../common/3-set-up-dns.sh
