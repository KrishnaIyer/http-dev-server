#!/bin/bash

set -e

chart=$1
keyname=$2

echo "Sign helm chart package $chart"
shasum=$(openssl dgst -sha256 $chart | awk '{ print $2 }')
echo $shasum
chartyaml=$(tar -zxf $chart --exclude 'charts/' -O '*/Chart.yaml')
c=$(cat << EOF
$chartyaml

...
files:
  $chart: sha256:$shasum
EOF
)
keyuser=""
if [ "$keyname" != "" ]; then
keyuser="-u $keyname"
fi
echo "$c" | gpg --clearsign -o "$chart.prov" $keyuser
