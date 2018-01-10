#!/bin/sh
set -x -e

cd $(dirname $0)

../common/0-clean-up.sh

../common/1-generate-config-json.sh
# TODO: create config from the docker image ?

## 2. Create rootfs
if [ -z $1 ]; then
    echo "Usage: $0 DOCKERIMAGE"
    exit 1
fi
image=$1
cid=$(docker create $image true)
docker export $cid | tar Cx ./rootfs || true
docker rm -f $cid

../common/3-set-up-dns.sh
