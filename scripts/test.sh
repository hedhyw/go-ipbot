#!/bin/bash

env=$1

source $env > /dev/null
cd "$(dirname "$0")/.." > /dev/null
vgo test ./... -count=1 -covermode=count -v