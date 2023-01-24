#!/bin/bash -eu

CJENV="/var/repositories/CJ-ENV"

useradd -s /bin/bash -d $CJENV/jenkins/ jenkins
