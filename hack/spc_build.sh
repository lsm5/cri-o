#!/bin/bash

set -e

if [[ "$SPC" != "true" ]]
then
    echo "This script is intended to be executed in an SPC,"
    echo "by run_ci_tests.sh. Using it otherwise may result"
    echo "in unplesent side-effects."
    exit 1
fi

echo
echo "Build Environment:"
env

set +x
apt-get -qq update
apt-get -qq install python-software-properties software-properties-common
add-apt-repository -y ppa:gophers/archive
apt-get -qq update
apt-get -qq install debhelper libdevmapper-dev libgpgme11-dev libseccomp-dev golang-1.10 btrfs-tools libglib2.0-dev go-md2man dh-golang
dpkg-buildpackage -us -uc
