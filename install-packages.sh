#!/bin/bash -eu

apt update
xargs -rd'\n' apt install -y <"packages.txt" | envsubst
apt upgrade -y
