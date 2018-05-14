#!/usr/bin/env bash
sudo nohup docker run \
    -v /home/nelsontk/Documents/dev/motion-vue/motion:/srv \
    -v /etc/config.json:/config.json \
    -p 80:80 \
    hacdias/filebrowser &
pid=$!

npm run serve
kill ${pid}
trap "echo kill ${pid}; exit 1;" INT
