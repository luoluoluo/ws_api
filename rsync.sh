#!/bin/bash
root=$(cd "$(dirname "$0")"; pwd)
rsync -avz -e ssh $root/dist/ l@wanshi.org:/data/ws_api
