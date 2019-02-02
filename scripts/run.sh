#!/bin/bash

env=$1
main=$2

source $env > /dev/null
cd "$(dirname "$0")/.." > /dev/null
vgo run $main