#!/bin/sh
set -e -x
cd $(dirname $0)
docker build -t runrootless-proot --target proot .
cid=$(docker create runrootless-proot none)
mkdir -p ~/.runrootless
docker cp ${cid}:/proot ~/.runrootless/runrootless-proot
docker rm -f ${cid}
