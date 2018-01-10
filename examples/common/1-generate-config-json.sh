#!/bin/sh
set -x -e

## 1. Generate config.json
runc spec
sed -i 's/"readonly": true/"readonly": false/g' config.json
