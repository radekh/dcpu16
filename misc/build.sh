#!/bin/bash
# filename: build.sh

# clone the repo
git clone -q "${1}" "clones/${2}"

cd "clones/${2}"

# update the submodules (how do we handle errors here?)
git submodule --quiet update --init --recursive

cd ../../

# tar up the directory
tar --exclude=.git -cf "archives/${2}-master.tar" "clones/${2}"

# remove the cloned repo
rm -rf "clones/${2}"

echo "A new archive has been created at /archives/${2}-master.tar" | mail -s "Github project build successful!" "radek.hnilica@gmail.com"