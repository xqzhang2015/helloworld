#!/usr/bin/env bash

# set -x

# example for 7GB: cleanup /ads/run/http_access 7000000000
cleanup() {
    local dir=${1}
    local threshold=${2:-10000000000}

    # echo ${dir}
    # echo ${threshold}

    cd ${dir}
    mysize=$(echo $(du -sb .) | awk '{print $1}')
    if (( ${mysize} > ${threshold} )); then rm -f $(ls -t | tail -1); fi
    cd -
}