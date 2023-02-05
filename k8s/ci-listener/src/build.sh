#!/bin/bash -eu

go build -ldflags="-extldflags=-static" -tags osusergo,netgo -o "ci-listener" && strip ci-listener
