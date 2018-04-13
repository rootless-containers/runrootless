#!/bin/sh
set -e

cd $(dirname $0)

if [ -z $1 ]; then
    /bin/echo -e "Usage: $0 SKOPEOIMAGE\ne.g. docker://ubuntu"
    exit 1
fi
image=$1

set -x

rm -rf skopeo-tmp umoci-bundle
skopeo copy ${image} oci:skopeo-tmp:tmp
umoci unpack --rootless --image skopeo-tmp:tmp umoci-bundle
rm -rf skopeo-tmp

echo "Generated bundle at ./umoci-bundle"
