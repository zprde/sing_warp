#!/bin/sh

export http_proxy=http://localhost:1080
export https_proxy=http://localhost:1080
wget -O /dev/null https://www.google.com/ncr
