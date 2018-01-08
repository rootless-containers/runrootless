#!/bin/sh
set -e -x
cd $(dirname $0)
docker build -t runrootless-proot .
cid=$(docker create runrootless-proot none)
mkdir -p ~/.runrootless
docker cp ${cid}:/runrootless-proot ~/.runrootless
docker rm -f ${cid}

